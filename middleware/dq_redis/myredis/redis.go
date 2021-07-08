package myredis

import (
	"log"
	"time"

	"github.com/garyburd/redigo/redis"
)

var (
	addr  = "127.0.0.1:6379"
	db    = 0
	count = 5
	// count = 10000
	pwd = "123"
)

func New() (res *redis.Pool) {
	// 建立连接池
	res = &redis.Pool{
		MaxIdle:     100,             //定义连接池中最大连接数（超过这个数会关闭老的链接，总会保持这个数）
		MaxActive:   20,              //最大的激活连接数，表示同时最多有N个连接
		IdleTimeout: 5 * time.Second, //定义链接的超时时间，每次p.Get()的时候会检测这个连接是否超时（超时会关闭，并释放可用连接数）
		Wait:        true,            // 当可用连接数为0是，那么当wait=true,那么当调用p.Get()时，会阻塞等待，否则，返回nil.
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", addr,
				redis.DialDatabase(db), //默认是索引0，可以自己指定
				// redis.DialPassword(pwd), //默认是空
				redis.DialReadTimeout(1*time.Second),
				redis.DialWriteTimeout(1*time.Second),
				redis.DialConnectTimeout(200*time.Millisecond))
			if err != nil {
				log.Fatalf("redis.Dial error(%v)\n", err)
				return nil, err
			}
			// if _, err := c.Do("AUTH", pwd); err != nil {
			// 	c.Close()
			// 	log.Fatalf("c.Do auth error(%v)\n", err)
			// 	return nil, err
			// }
			if _, err := c.Do("SELECT", db); err != nil { //选择db
				c.Close()
				return nil, err
			}
			return c, err
		},

		// 如果设置了给func,那么每次p.Get()的时候都会调用改方法来验证连接的可用性
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
	}

	return
}
