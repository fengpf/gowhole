package dq_redis

import (
	"github.com/google/uuid"
	"gowhole/middleware/log"
	"sync"
	"testing"
	"time"
)

const _testTopic = "topic-task-test"

func Test_dq(t *testing.T) {
	dq := NewDQ(&Config{
		BucketCount: 5,
	})
	dq.Start()
	//defer dq.Close()

	i := 20
	for i > 0 {
		job := Job{
			Topic: _testTopic,
			ID:    int64(uuid.New().ID()),
			Delay: int64(i),
			TTR:   3,
		}

		err := dq.EnQueue(job)
		if err != nil {
			log.Error("Put job(%+v) error(%v)", job, err)
		}

		i--

	}

	log.Info("enqueue fin")

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()

		for {
			job, err := dq.DeQueue(_testTopic)

			if err != nil {
				//t.Errorf("Pop error(%v)\n", err)
				time.Sleep(time.Second)
				continue
			}

			if job == nil {
				time.Sleep(time.Second)
				continue
			}

			log.Info("业务处理:job%+v now(%s) expire(%s)",
				job,
				time.Now().Format("2006-01-02 15:04:05"),
				time.Unix(job.TimeStamp, 0).Format("2006-01-02 15:04:05"),
			)

		}
	}()
	wg.Wait()
}
