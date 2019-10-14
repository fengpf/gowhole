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
		dl     = redis_instance.New("EX")

		//mux sync.Mutex
	)

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			g := "goroutine-" + strconv.Itoa(i)

			//mux.Lock()
			dl.SetConn(redis_instance.RedisClient.Get())
			res, err := dl.Lock(key, expire)
			if err != nil {
				return
			}
			defer func() {
				dl.UnLock(key)
				//mux.Lock()
			}()

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
