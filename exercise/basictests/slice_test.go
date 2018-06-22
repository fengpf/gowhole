package basictests

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"testing"
	"time"

	"github.com/davecgh/go-spew/spew"

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

type pagination struct {
	Pn    int
	Ps    int
	Total int
	Items []int
}

func Test_pagination(t *testing.T) {
	var (
		pn, ps int
	)
	items := []int{1, 2, 3, 5}

	pn = 4
	ps = 1

	res := &pagination{}
	total := len(items)
	res.Total = total
	start := (pn - 1) * ps
	end := pn * ps

	println(total, start, end)
	if end < total {
		println(111)
		res.Items = items[start:end]
	} else {
		println(222)
		res.Items = items[start:total]
	}

	spew.Dump(res)
}

func Test_cutOut(t *testing.T) {
	/* 创建切片 */
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	printSlice(numbers)

	/* 打印原始切片 */
	fmt.Println("numbers ==", numbers)

	/* 打印子切片从索引1(包含) 到索引4(不包含)*/
	fmt.Println("numbers[1:4] ==", numbers[1:4])

	/* 默认下限为 0*/
	fmt.Println("numbers[:3] ==", numbers[:3])

	/* 默认上限为 len(s)*/
	fmt.Println("numbers[4:] ==", numbers[4:])

	numbers1 := make([]int, 0, 5)
	printSlice(numbers1)

	/* 打印子切片从索引  0(包含) 到索引 2(不包含) */
	number2 := numbers[:2]
	printSlice(number2)

	/* 打印子切片从索引 2(包含) 到索引 5(不包含) */
	number3 := numbers[2:5]
	printSlice(number3)

}

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}

func Test_make(t *testing.T) {
	arcs := []int{
		11, 12, 13,
	}
	var bindMapTIDs map[int][]int
	bindMapTIDs = make(map[int][]int)
	for k, v := range arcs {
		bindMapTIDs[k] = append(bindMapTIDs[k], v)
	}
	fmt.Println(bindMapTIDs)
}

func CheckSameElementOfSlice(a []int64, b []int64) (isSame bool) {
	isSame = true
	if len(a) == 0 && len(b) == 0 {
		return
	}
	isHas := make(map[int64]bool)
	if len(a) < len(b) {
		for _, k := range a {
			isHas[k] = true
		}
		for _, v := range b {
			if !isHas[v] {
				isSame = false
			}
		}
	} else {
		for _, k := range b {
			isHas[k] = true
		}
		for _, v := range a {
			if !isHas[v] {
				isSame = false
			}
		}
	}
	return
}

func Test_checkSameElementOfSlice(t *testing.T) {
	a := []int64{9, 10, 11, 12}
	b := []int64{16, 2, 3, 15}
	c := []int64{2, 3, 9, 10, 11, 12}
	// fmt.Println(CheckSameElementOfSlice(a, b))
	fmt.Println(ContainAll(a, c))
	fmt.Println(ContainAtLeastOne(b, c))
}

func ContainAll(a []int64, b []int64) bool {
	isHas := make(map[int64]bool)
	for _, k := range b {
		isHas[k] = true
	}
	for _, v := range a {
		if !isHas[v] {
			return false
		}
	}
	return true
}

func ContainAtLeastOne(a []int64, b []int64) bool {
	isHas := make(map[int64]bool)
	for _, k := range b {
		isHas[k] = true
	}
	for _, v := range a {
		if isHas[v] {
			return true
		}
	}
	return false
}
