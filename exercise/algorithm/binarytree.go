package main

import "fmt"

type Node struct {
	Parent *Node
	Left   *Node
	Right  *Node
}

type Tree struct {
	Root *Node
}

func Eval(tree Tree) int {
	root := tree.Root
	i := 0
	for root.Left == nil || root.Right == nil {
		if root.Left == nil && root.Right != nil {
			root = root.Right
		} else if root.Left != nil && root.Right == nil {
			root = root.Left
		} else {
			return i
		}
		i++
	}
	l := depth(root.Left)
	r := depth(root.Right)
	result1 := l + r + 2
	result2 := l + i + 1
	result3 := r + i + 1
	if result1 >= result2 && result1 >= result3 {
		return result1
	}
	if result2 >= result1 && result2 >= result3 {
		return result2
	}
	return result3
}

func depth(node *Node) int {
	var helper func(*Node, int) int
	helper = func(node *Node, d int) int {
		var l, r int
		if node.Left != nil {
			l = helper(node.Left, d+1)
		}
		if node.Right != nil {
			r = helper(node.Right, d+1)
		}
		if node.Left == nil && node.Right == nil {
			return d
		} else {
			if l > r {
				return l
			} else {
				return r
			}
		}
	}
	return helper(node, 0)
}

func main() {
	/*
		tree := Tree{
			Root: &Node{
				Left: &Node{
					Left: &Node{
						Left: &Node{
							Right: &Node{},
						},
					},
					Right: &Node{
						Left: &Node{
							Left: &Node{
								Left: &Node{
									Left: &Node{},
								},
							},
						},
					},
				},
				Right: &Node{
					Right: &Node{
						Left:  &Node{},
						Right: &Node{},
					},
				},
			},
		}
	*/
	tree := Tree{
		Root: &Node{
			Left: &Node{
				Right: &Node{
					Left:  &Node{},
					Right: &Node{},
				},
			},
		},
	}
	fmt.Println(Eval(tree))
}
