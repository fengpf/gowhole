package main

import (
	"context"
	"fmt"

	"github.com/tsuna/gohbase"
	"github.com/tsuna/gohbase/hrpc"
)

func init() {
	// 以Stdout为输出，代替默认的stderr
	// logrus.SetOutput(os.Stdout)
	// 设置日志等级
	// logrus.SetLevel(logrus.DebugLevel)
}

func main() {
	client := gohbase.NewClient("172.18.33.75")
	//Get an entire row
	getRequest, err := hrpc.NewGetStr(context.Background(), "table", "row")
	if err != nil {
		fmt.Println(err)
		return
	}
	getRsp, err := client.Get(getRequest)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(getRsp)
}
