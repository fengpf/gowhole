package main

import (
	"log"
	"strconv"
	"sync"

	"gowhole/middleware/redis/lock/redis_instance"
)

func main() {
	ConcurrentCount()
}

func ConcurrentCount() {
	var (
		count  = 0
		key    = "test"
		expire = 10
	)

	var wg sync.WaitGroup
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			g := "goroutine-" + strconv.Itoa(i)
			dl := redis_instance.New(redis_instance.RedisClient.Get(), "EX")
			res, err := dl.Lock(key, expire)
			defer dl.UnLock(key)
			if err != nil {
				return
			}

			if !res { //没拿到锁，返回
				log.Println(g + " miss get lock")
				return
			}

			log.Println(g + " get Lock ")
			count++
			log.Println(g + " " + strconv.Itoa(count))

		}(i)
	}
	wg.Wait()
}
