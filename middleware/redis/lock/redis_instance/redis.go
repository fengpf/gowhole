package redis_instance

import (
	"fmt"
	"log"
	"time"

	"github.com/garyburd/redigo/redis"
)

var (
	RedisClient *redis.Pool
	host        = "127.0.0.1:6379"
)

func init() {
	// 建立连接池
	RedisClient = &redis.Pool{
		MaxIdle:     20,                //最大的空闲连接数，表示即使没有redis连接时依然可以保持N个空闲的连接，而不被清除，随时处于待命状态。
		MaxActive:   100,               //最大的激活连接数，表示同时最多有N个连接
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

type DisLock struct {
	conn redis.Conn
	unit string //PX milliseconds,EX seconds
}

func New(c redis.Conn, unit string) *DisLock {
	return &DisLock{
		conn: c,
		unit: unit,
	}
}

func (d *DisLock) Close(c redis.Conn) {
	d.conn.Close()
}

//SET key value [EX seconds] [PX milliseconds] [NX|XX]

//可选参数
//EX seconds：将键的过期时间设置为seconds秒。 执行SET key value EX seconds的效果等同于执行SETEX key seconds value。
//PX milliseconds：将键的过期时间设置为milliseconds毫秒。 执行SET key value PX milliseconds的效果等同于执行 PSETEX key milliseconds value。
//NX ： 只在键不存在时， 才对键进行设置操作。 执行SET key value NX的效果等同于执行SETNX key value。
//XX ： 只在键已经存在时， 才对键进行设置操作。
//Lock .

//在Redis 2.6.12版本以前，SET命令总是返回 OK 。
//从Redis 2.6.12版本开始，SET命令只在设置操作成功完成时才返回OK；
//如果命令使用了NX或者XX选项， 但是因为条件没达到而造成设置操作未执行， 那么命令将返回空批量回复（NULL Bulk Reply）

func (d *DisLock) Lock(key string, expire int) (res bool, err error) {
	var reply interface{}
	reply, err = d.conn.Do("SET", key, "1", d.unit, expire, "NX")
	if err != nil {
		log.Printf("setnx key(%s) unit(%s) error(%v)", key, d.unit, err.Error())
		return
	}

	println(reply) //0x1237a80,0x145c020

	log.Printf("%v", reply)

	if reply != nil {
		res = true
	}
	return
}

//UnLock .
func (d *DisLock) UnLock(key string) (err error) {
	val, err := d.conn.Do("GET", key)
	if err != nil {
		log.Printf("get key(%s) error(%v)", key, err.Error())
		return
	}

	if val == key {
		_, err = d.conn.Do("DEL", key)
	}
	return
}
