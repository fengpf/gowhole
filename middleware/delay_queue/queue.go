package main

import (
	"container/list"
	"sync"
)

type Queue struct {
	l         *list.List
	mux       sync.Mutex
	cond sync.Cond
}

func New() *Queue {
	return &Queue{
		l: list.New(),
	}
}

func (q *Queue) offer(e interface{}) bool {
	q.mux.Lock()
	defer q.mux.Unlock()

	q.l.PushFront(e) //放一个元素
	if q.l.Front().Value == e {
		q.cond.Broadcast()
	}

	return true
}

func (q *Queue) poll() interface{} {
	q.mux.Lock()
	defer q.mux.Unlock()

	if q.l.Front().Value == nil { //时间不到不走
		return nil
	} else {
		return q.poll() //  取出优先级队列里面第0个元素，优先级队列是根据时间排序了的
	}
}

func (q *Queue) BlockPop() interface{} {
	q.mux.Lock()
	defer q.mux.Unlock()

	for {
		first:=q.l.Front().Value

		// 队首为空，则阻塞当前线程
		if first == nil {
			q.cond.Wait() //等待offer方法唤醒，再次循环去peek，
		} else {

			var delay int64
			// 获取队首元素的超时时间
			//delay := first.getDelay(NANOSECONDS);//返回还需等待3秒

			// 已超时，直接出队
			if delay <= 0 {
				return q.poll();
			}

			// 释放first的引用，避免内存泄漏
			first = nil

			//有元素，但是元素没到时间

		}
	}
}
