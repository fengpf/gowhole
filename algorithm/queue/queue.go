package main

import "sync"

const (
	bucketSize  = 2
	maxfreelist = 3
)

var usefreelist = true

type Queue struct {
	sync.Mutex
	count    uint64
	head     *bucket
	tail     *bucket
	numfree  int
	freelist *bucket
}

type bucket struct {
	data [bucketSize]int
	head byte
	tail byte
	next *bucket
}

func (q *Queue) Enqueue(v int) {
	q.Lock()

	b := q.tail
	if b == nil || b.tail == bucketSize { //b.tail == bucketSize 如果当前桶满则新new
		b = nil
		if usefreelist && q.numfree > 0 {
			b = q.freelist
			q.freelist = b.next
			q.numfree--

			b.head = 0
			b.tail = 0
			b.next = nil
		}

		if b == nil {
			b = new(bucket)
		}

		if q.head == nil {
			q.head = b
		}

		if q.tail == nil {
			q.tail = b
		} else {
			q.tail.next = b
			q.tail = b
		}
	}

	b.data[b.tail] = v
	b.tail++
	q.count++

	q.Unlock()
}

func (q *Queue) Dequeue() (v int, ok bool) {
	q.Lock()
	if q.count == 0 {
		q.Unlock()
		return 0, false
	}

	v = q.head.data[q.head.head]
	q.head.head++
	q.count--

	if q.head.head == bucketSize || q.count == 0 {
		x := q.head     //使用x存队列头节点
		q.head = x.next //q的head 指向x的next

		if q.count == 0 { //如果队列长度是0，则x取q的尾部节点，清空队列的头尾节点
			x = q.tail
			q.head = nil
			q.tail = nil
		}

		if usefreelist && q.numfree < maxfreelist { //如果开启使用回收队列，并且回收队列的数量小于最大释放的队列长度
			x.next = q.freelist //x的next 指向q的回收队列
			q.freelist = x      //将x添加到回收队列
			q.numfree++         //回收队列长度增加
		}
	}

	q.Unlock()
	ok = true
	return
}
