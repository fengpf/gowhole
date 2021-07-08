package main

import (
	"fmt"
	"gowhole/middleware/redis/lock/redis_instance"
	"log"
	"strconv"
	"sync"

	"github.com/garyburd/redigo/redis"
	uuid "github.com/satori/go.uuid"
)

func getLock(conn redis.Conn, key string, timeout int64) (value string, err error) {
	u := uuid.NewV4()

	value = u.String()
	_, err = redis.String(conn.Do("SET", key, value, "EX", timeout, "NX"))
	if err == redis.ErrNil {
		return "", nil
	}
	if err != nil {
		log.Fatalf("conn.Do set nx:error(%v)", err)
		return "", err
	}
	return value, nil

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

//独立增加过期时间
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
			name := "thread-" + strconv.Itoa(i)

			conn := redis_instance.RedisClient.Get()
			value, err := getLock(conn, key, int64(DefaultExpire))
			if err != nil {
				log.Fatalf("Error while attempting lock error(%v)", err)
				return
			}

			if value == "" {
				println(name + " do not get Lock ")
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
