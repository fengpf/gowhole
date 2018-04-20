package base

import (
	"encoding/json"
	"fmt"
	"go-common/log"
	"os"
	"testing"
	"time"
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
	b, err := json.Marshal(res)
	if err != nil {
		fmt.Printf("error(%v)", err)
	}
	os.Stdout.Write(b)
}
