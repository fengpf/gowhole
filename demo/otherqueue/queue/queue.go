package queue

import (
	"sync"
	"time"
)

type Queue struct {
	mutex   sync.Mutex
	storage []interface{}
	length  int
}

func NewQueue() *Queue {
	return &Queue{
		storage: []interface{}{},
		length:  0,
	}
}

func (q *Queue) Put(val interface{}) {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	q.storage = append(q.storage, val)
	q.length++
}

func (q *Queue) Get() (interface{}, bool) {
	q.mutex.Lock()
	defer q.mutex.Unlock()

	if q.length == 0 {
		return nil, false
	}

	out := q.storage[0]
	q.storage = q.storage[1:]
	q.length--
	return out, true
}

func (q *Queue) Show() []interface{} {
	q.mutex.Lock()
	defer q.mutex.Unlock()

	res := make([]interface{}, len(q.storage))
	copy(res, q.storage)
	return res
}

func (q *Queue) Delete(f func(elem interface{}) bool) {
	q.mutex.Lock()
	defer q.mutex.Unlock()

	res := make([]interface{}, len(q.storage))
	copy(res, q.storage)

	for i := len(res) - 1; i >= 0; i-- {
		if f(res[i]) {
			q.storage = append(q.storage[:i], q.storage[i+1:]...)
		}
	}
}

type DelayCache struct {
	storage []*DelayTask
	mutex   sync.Mutex
}

type DelayTask struct {
	Task      interface{}
	notBefore time.Time
}

func NewDelayCache() *DelayCache {
	return &DelayCache{
		storage: []*DelayTask{},
	}
}

// delay/second
func (d *DelayCache) AddDelayTask(task interface{}, delay int) {
	d.mutex.Lock()
	defer d.mutex.Unlock()

	d.storage = append(d.storage, &DelayTask{
		Task:      task,
		notBefore: time.Now().Add(time.Duration(delay) * time.Second),
	})
}

func (d *DelayCache) GetDelayTasks(t time.Time) (tasks []interface{}) {
	d.mutex.Lock()
	defer d.mutex.Unlock()

	tasks = []interface{}{}
	cp := make([]*DelayTask, len(d.storage))
	copy(cp, d.storage)

	for i := len(cp) - 1; i >= 0; i-- {
		if cp[i].notBefore.Before(t) {
			tasks = append(tasks, cp[i].Task)
			d.storage = append(d.storage[:i], d.storage[i+1:]...)
		}
	}

	return
}
