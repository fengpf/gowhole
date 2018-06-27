package binarytree

import (
	"flag"
	"fmt"
	"runtime"
	"strconv"
	"sync"
	"testing"
)

var (
	minDepth = 4
	n        = 0
)

func bottomUpTree(depth int) *Node {
	if depth <= 0 {
		return &Node{nil, nil}
	}
	return &Node{
		bottomUpTree(depth - 1),
		bottomUpTree(depth - 1),
	}
}

type Node struct {
	left, right *Node
}

func (n *Node) ItemCheck() int {
	if n.left == nil {
		return 1
	}
	return 1 + n.left.ItemCheck() + n.right.ItemCheck()
}

func Test_bTree(t *testing.T) {
	runtime.GOMAXPROCS(runtime.NumCPU() * 2)

	flag.Parse()
	if flag.NArg() > 0 {
		n, _ = strconv.Atoi(flag.Arg(0))
	}

	maxDepth := n
	if minDepth+2 > n {
		maxDepth = minDepth + 2
	}
	stretchDepth := maxDepth + 1

	checkL := bottomUpTree(stretchDepth).ItemCheck()
	fmt.Printf("stretch tree of depth %d\t check: %d\n", stretchDepth, checkL)

	longLivedTree := bottomUpTree(maxDepth)

	resultTrees := make([]int, maxDepth+1)
	resultCheck := make([]int, maxDepth+1)

	var wg sync.WaitGroup
	for depthL := minDepth; depthL <= maxDepth; depthL += 2 {
		wg.Add(1)
		go func(depth int) {
			iterations := 1 << uint(maxDepth-depth+minDepth)
			check := 0

			for i := 1; i <= iterations; i++ {
				check += bottomUpTree(depth).ItemCheck()
			}
			resultTrees[depth] = iterations
			resultCheck[depth] = check

			wg.Done()
		}(depthL)
	}
	wg.Wait()
	for depth := minDepth; depth <= maxDepth; depth += 2 {
		fmt.Printf("%d\t trees of depth %d\t check: %d\n",
			resultTrees[depth], depth, resultCheck[depth],
		)
	}
	fmt.Printf("long lived tree of depth %d\t check: %d\n",
		maxDepth, longLivedTree.ItemCheck(),
	)
}
