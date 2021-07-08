package main

// This example demonstrates a priority queue built using the heap interface.

import (
	"container/heap"
	"time"
)

type Delayed interface {
	GetDelay() time.Duration
}

// An Item is something we manage in a priority queue.
type Item struct {
	value    string // The value of the item; arbitrary.
	priority int64  // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
	// when task is ready, execute OnTrigger function
	OnTrigger func(time.Time, string)
}

// triggerTime is time of the task should be execute
func NewDelayedItemFunc(triggerTime time.Time, value string, f func(time.Time, string)) *Item {
	item := Item{}
	item.priority = triggerTime.UnixNano()
	item.value = value
	item.OnTrigger = f
	return &item
}

func (item *Item) GetDelay() time.Duration {
	return time.Duration(item.priority - time.Now().UnixNano())
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq *PriorityQueue) Len() int { return len(*pq) }

func (pq *PriorityQueue) Less(i, j int) bool {
	return (*pq)[i].priority < (*pq)[j].priority
}

func (pq *PriorityQueue) Swap(i, j int) {
	(*pq)[i], (*pq)[j] = (*pq)[j], (*pq)[i]
	(*pq)[i].index = i
	(*pq)[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *Item, value string, priority int64) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}

func (pq *PriorityQueue) Peek() interface{} {
	old := *pq
	n := len(old)
	if n <= 0 {
		return nil
	}
	return old[0]
}

func (pq *PriorityQueue) Clear() {
	pq = nil
}
