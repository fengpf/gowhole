package test

import (
	"fmt"
	"os"
	"sort"
	"sync"
	"testing"

	"encoding/json"
)

type stu struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func Test_share(t *testing.T) {
	t.Parallel()
	ids := make([]int, 0, 8)
	for i := 0; i < 8; i++ {
		ids = append(ids, i)
		// if len(ids) == 3 {
		// 	break
		// }
	}
	// fmt.Println(ids)
	reverse(ids)
	// fmt.Println(ids)
	sort.Slice(ids, func(i, j int) bool {
		return ids[i] < ids[j]
	})
	var (
		wg sync.WaitGroup
		mx sync.RWMutex
	)
	wg.Add(len(ids))
	stus := make([]*stu, 0, len(ids))
	stusMap := make(map[int]*stu, len(ids))
	var (
	// s *stu //多个goroutine修改共享变量
	)
	for _, id := range ids {
		go func(i int) {
			defer wg.Done()
			s := getName(i)
			mx.Lock()
			stusMap[i] = s
			mx.Unlock()
		}(id)
	}
	wg.Wait()
	// fmt.Println(stusMap)
	for _, id := range ids {
		if s, ok := stusMap[id]; ok {
			stus = append(stus, s)
		}
	}
	b, err := json.Marshal(stus)
	if err != nil {
		fmt.Printf("error(%v)", err)
	}
	os.Stdout.Write(b)
}

func getName(i int) (s *stu) {
	var n string
	switch i {
	case 0:
		n = "tom"
	case 1:
		n = "jack"
	case 2:
		n = "li"
	case 3:
		n = "33"
	case 4:
		n = "44"
	case 5:
		n = "55"
	case 6:
		n = "66"
	case 7:
		n = "77"
	case 8:
		n = "88"
	default:
		n = "default"
	}
	s = &stu{
		ID:   i,
		Name: n,
	}
	return
}
