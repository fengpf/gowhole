package main

import (
	"fmt"
	"gowhole/exercise/algorithm/tree/bplustree"
	"time"
)

func main() {
	bplustree_insert()
}

func bplustree_insert() {
	testCount := 1000000
	bt := bplustree.NewBTree()
	start := time.Now()
	for i := testCount; i > 0; i-- {
		bt.Insert(i, "")
	}
	fmt.Println(time.Now().Sub(start))
}
