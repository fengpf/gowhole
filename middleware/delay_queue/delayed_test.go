package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

func DefaultAction(t time.Time, value string) {
	fmt.Printf("trigger_time:%v, value:%v\n", t, value)
}

func push(timer *SuperTimer, name string) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 20; i++ {
		now := time.Now()
		t := now.Add(time.Duration(r.Int63n(100)) * time.Second)
		value := fmt.Sprintf("%v:value:%v", name, strconv.Itoa(i))

		item := NewDelayedItemFunc(t, value, DefaultAction)
		timer.Add(item)
	}
}

func Test_timer(t *testing.T) {
	timer := NewSuperTimer(5)
	for i := 0; i < 5; i++ {
		go push(timer, "worker"+strconv.Itoa(i))
	}
	time.Sleep(100)
}

func Test_Message(t *testing.T) {
	timer := NewSuperTimer(2)

	now := time.Now()
	tm := now.Add(time.Duration(3 * time.Second))
	tm2 := now.Add(time.Duration(2 * time.Second))
	tm3 := now.Add(time.Duration(1 * time.Second))

	timer.Add(NewDelayedItemFunc(tm, "world", DefaultAction))
	timer.Add(NewDelayedItemFunc(tm2, "hello", DefaultAction))
	timer.Add(NewDelayedItemFunc(tm3, "wo", DefaultAction))

	time.Sleep(10 * time.Second)
}
