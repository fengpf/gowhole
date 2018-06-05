package basictests

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"testing"
	"time"

	"github.com/qiniu/log"
)

func Test_Slice(t *testing.T) {
	a := make([]int, 0, 3)
	for i := 0; i < 5; i++ {
		a = append(a, i)
		if len(a) == 3 {
			break
		}
	}
	fmt.Println(a)
}

type ThirtyDay struct {
	DateKey   int64 `json:"date_key"`
	TotalIncr int64 `json:"total_inc"`
}

func Test_AppendByTime(t *testing.T) {
	res := make([]*ThirtyDay, 0, 5)
	td := &ThirtyDay{
		DateKey:   1523923200,
		TotalIncr: 12,
	}
	// td2 := &ThirtyDay{
	// 	DateKey:   1524009600,
	// 	TotalIncr: 6,
	// }
	res = append(res, td)
	// res = append(res, td2)
	timeSlice := make([]int64, 0, 5)

	resMap := make(map[int64]int64)

	for i := 5; i > 0; i-- {
		dt := time.Now().AddDate(0, 0, -1-i).Add(-12 * time.Hour).Format("20060102")
		tm, err := time.Parse("20060102", dt)
		if err != nil {
			log.Error("time.Parse error(%v)", err)
			return
		}
		t := tm.Unix()
		timeSlice = append(timeSlice, t)
	}
	for _, v := range res {
		resMap[v.DateKey] = v.DateKey
	}
	if len(res) < 5 {
		for _, v := range timeSlice {
			fmt.Println(v)
			if _, ok := resMap[v]; !ok {
				td := &ThirtyDay{
					DateKey: v,
				}
				res = append(res, td)
			}
			if len(res) == 5 {
				break
			}
		}
	}
	res = res[0:2]
	b, err := json.Marshal(res)
	if err != nil {
		fmt.Printf("error(%v)", err)
	}
	os.Stdout.Write(b)
}

type Person struct {
	Name string
	Age  int
}

func Test_sliceSort(t *testing.T) {
	a := []int{5, 2, 1, 4, 3}
	sort.Slice(a, func(i, j int) bool {
		return a[i] > a[j]
	})
	fmt.Println(a, sort.SliceIsSorted(a, func(i, j int) bool {
		return a[i] > a[j]
	}))

	var people = []Person{
		{"Bob", 31},
		{"AJohn", 42},
		{"DMichael", 17},
		{"CJenny", 26},
	}
	sort.Slice(people, func(i int, j int) bool {
		return people[i].Name > people[j].Name
	})

	fmt.Println(people)

}

// func BenchmarkSort(b *testing.B) {
// 	var people = []Person{
// 		{"Bob", 31},
// 		{"John", 42},
// 		{"Michael", 17},
// 		{"Jenny", 26},
// 	}
// 	for i := 0; i < b.N; i++ {
// 		sort.Sort(ByAge(people))
// 	}
// }

func BenchmarkSortSlice(b *testing.B) {
	var people = []Person{
		{"Bob", 31},
		{"John", 42},
		{"Michael", 17},
		{"Jenny", 26},
	}
	for i := 0; i < b.N; i++ {
		sort.Slice(people, func(i int, j int) bool {
			return people[i].Age < people[j].Age
		})
	}
}

func Test_deduplication(t *testing.T) {
	a := []int{2, 3, 5, 2, 4, 5, 6}
	b := make([]int, 0)
	fmt.Println(a)
	for k, v := range a {
		if len(b) == 0 {
			b = append(b, v)
		}
		fmt.Println("out cycle-----", k, v)
		for kk, vv := range b {
			if vv == v {
				fmt.Println("break-----", kk, v)
				break
			}
			fmt.Println("current-----", kk, b)
			if kk == len(b)-1 {
				b = append(b, v)
				fmt.Println("append----", kk, b)
			}
		}
	}
	fmt.Println(b)
}

func Test_copy(t *testing.T) {
	abortIndex := 1<<8 - 1
	Element := []int{1, 2, 3}
	ElementA := []int{4, 5, 6, 7}
	finalSize := len(Element) + len(ElementA)
	if finalSize >= int(abortIndex) {
		panic("too many Element")
	}
	mergedElement := make([]int, finalSize)
	copy(mergedElement, Element)
	fmt.Println(abortIndex, mergedElement)
	copy(mergedElement[len(Element):], ElementA)
	fmt.Println(mergedElement)
}
