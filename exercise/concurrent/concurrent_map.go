package main

import "sync"
import "fmt"


type SafeMap struct {
    sync.RWMutex
    Map map[int]int
}

func main() {
    const workers = 2
    var wg sync.WaitGroup
    wg.Add(workers)

    safeMap := new(SafeMap)
    safeMap.Map = make(map[int]int)

    for i := 0; i <= workers; i++ {
        go func(i int) {
            for j := 0; j < i; j++ {
                safeMap.writeMap(i,i)
                safeMap.readMap(i)
            }
            wg.Done()
        }(i)
    }
    wg.Wait()
}

func (sm *SafeMap) readMap(key int) int {
    sm.RLock()
    value := sm.Map[key]
    sm.RUnlock()
    fmt.Println("key=",key,"value=",value)
    return value
}

func (sm *SafeMap) writeMap(key int, value int) {
    sm.Lock()
    sm.Map[key] = value
    sm.Unlock()
}