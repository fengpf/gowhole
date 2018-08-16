package main

import (
	"fmt"
	"log"
	"strconv"
	"sync"
	"time"

	"github.com/garyburd/redigo/redis"
	uuid "github.com/satori/go.uuid"
)

var (
	redisClient *redis.Pool
	host        = "127.0.0.1:6379"
)

func init() {
	// 建立连接池
	redisClient = &redis.Pool{
		MaxIdle:     5,                 //最大的空闲连接数，表示即使没有redis连接时依然可以保持N个空闲的连接，而不被清除，随时处于待命状态。
		MaxActive:   20,                //最大的激活连接数，表示同时最多有N个连接
		IdleTimeout: 180 * time.Second, //最大的空闲连接等待时间，超过此时间后，空闲连接将被关闭
		Dial: func() (redis.Conn, error) { //建立连接
			c, err := redis.Dial("tcp", host)
			if err != nil {
				fmt.Printf("redis.Dial error(%v)\n", err)
				return nil, err
			}
			// if _, err := c.Do("AUTH", pwd); err != nil {
			// 	c.Close()
			// 	fmt.Printf("c.Do auth error(%v)\n", err)
			// 	return nil, err
			// }
			// c.Do("SELECT", 0)//选择db
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
	}
}

func getLock(conn redis.Conn, key string, timeout int64) (value string, err error) {
	timeWait := time.Now().Unix() + timeout*1000
	u, err := uuid.NewV4()
	if err != nil {
		fmt.Printf("Something gone wrong: %s", err)
		return
	}
	value = u.String()
	for time.Now().Unix() < timeWait {
		_, err = redis.String(conn.Do("SET", key, value, "EX", timeout, "NX"))
		if err == redis.ErrNil {
			return "", nil
		}
		if err != nil {
			log.Fatalf("conn.Do set nx:error(%v)", err)
			return "", err
		}
		return value, nil
	}
	fmt.Println("get lock timeouts")
	return "", nil
}

func unlock(conn redis.Conn, key, value string) (err error) {
	val, err := conn.Do("GET", key)
	if err != nil {
		return
	}
	if val == value {
		_, err = conn.Do("DEL", key)
	}
	return
}

func addExpire(conn redis.Conn, key, value string, exTime int64) (ok bool, err error) {
	//当 key 不存在时，返回 -2 。 当 key 存在但没有设置剩余生存时间时，返回 -1 。 否则，以毫秒为单位，返回 key 的剩余生存时间。
	ttlTime, err := redis.Int64(conn.Do("TTL", key))
	if err != nil {
		log.Fatalf("redis get failed:error(%v)", err)
		return
	}
	if ttlTime > 0 { //如果之前key存在过期时间，则在此基础上增加过期时间
		_, err := redis.String(conn.Do("SET", key, value, "EX", int(ttlTime+exTime)))
		if err == redis.ErrNil {
			return false, nil
		}
		if err != nil {
			log.Fatalf("AddTimeout lock.conn.Do:error(%v)", err)
			return false, err
		}
	}
	return false, nil
}

func main() {
	DefaultExpire := 100
	count := 0
	key := "test"
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			conn := redisClient.Get()
			name := "thread-" + strconv.Itoa(i)
			value, err := getLock(conn, key, int64(DefaultExpire))
			if err != nil {
				log.Fatalf("Error while attempting lock error(%v)", err)
				return
			}
			if value == "" {
				println(name + "do not get Lock ")
				return
			}
			println(name + " get Lock " + value)
			count++
			unlock(conn, key, value)
			println(name + " " + strconv.Itoa(count))
		}(i)
	}
	wg.Wait()
}
