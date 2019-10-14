package redis_instance

import (
	"log"
	"testing"
)

func Test_Lock(t *testing.T) {
	var (
		key    = "aaa"
		expire = 20
		dl     = New(RedisClient.Get(), "EX")
	)

	res, err := dl.Lock(key, expire)
	defer dl.UnLock(key)
	if err != nil {
		return
	}

	if !res { //没拿到锁，返回
		log.Println("miss get lock")
		return
	}

	log.Println("get lock")
}
