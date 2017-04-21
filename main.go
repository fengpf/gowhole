package main

import (
	"fmt"
	"gostudy/imgprocessing"
)

func main() {
	err := imgprocessing.DrawWaterMark("golang水印")
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	fmt.Printf("%s\n", "create watermark success ...")
}
