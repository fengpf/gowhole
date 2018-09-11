package main

import (
	"fmt"
)

type Node struct {
	Next  *Node
	Value interface{}
}

func main() {

	nums := []int{1, 2, 3, 4, 5}

	head := &Node{ //链表头指针
		Value: nums[0],
	}
	tail := head                     //链表尾部和头部公用一个指针节点
	for i := 1; i < len(nums); i++ { //构建指针链表
		tail.Next = &Node{
			Value: nums[i],
		}
		tail = tail.Next
	}

	tail.Next = head //尾部节点指向头节点，使得链表成环

	i := 0
	for n := head; n != nil; n = n.Next {
		fmt.Println("cycle linklist print ", n.Value)
		i++

		if i == 15 {
			break
		}
	}

	visited := make(map[*Node]bool)
	// visited[head.Next] = true
	// visited[tail] = true

	for n := head; n != nil; n = n.Next {
		if visited[n] {
			fmt.Println("cycle detected node", n.Value)
			break
		}
		fmt.Println("not visited node print", n.Value)
		visited[n] = true
	}
}
