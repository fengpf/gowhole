package synchronized

import (
	"fmt"
	"sync"
	"testing"
)

//安全的Map
type SynchronizedMap struct {
	rw   *sync.RWMutex
	data map[interface{}]interface{}
}

//存储操作
func (sm *SynchronizedMap) Put(k, v interface{}) {
	sm.rw.Lock()
	defer sm.rw.Unlock()
	sm.data[k] = v
}

//获取操作
func (sm *SynchronizedMap) Get(k interface{}) interface{} {
	sm.rw.RLock()
	defer sm.rw.RUnlock()
	return sm.data[k]
}

//删除操作
func (sm *SynchronizedMap) Delete(k interface{}) {
	sm.rw.Lock()
	defer sm.rw.Unlock()
	delete(sm.data, k)
}

//遍历Map，并且把遍历的值给回调函数，可以让调用者控制做任何事情
func (sm *SynchronizedMap) Each(cb func(interface{}, interface{})) {
	sm.rw.RLock()
	defer sm.rw.RUnlock()
	for k, v := range sm.data {
		cb(k, v)
	}
}

//生成初始化一个SynchronizedMap
func NewSynchronizedMap() *SynchronizedMap {
	return &SynchronizedMap{
		rw:   new(sync.RWMutex),
		data: make(map[interface{}]interface{}),
	}
}

func (sm *SynchronizedMap) printMap(k interface{}, v interface{}) {
	fmt.Println(k, " is ", v)
}

func Test_sync(t *testing.T) {
	sm := NewSynchronizedMap()
	sm.Put("name", "luby")
	sm.Put("age", 10)
	sm.Each(sm.printMap)
}
