package engine

import (
	"strings"
)

type node struct {
	pattern  string  //待匹配的路由， 例如 /p/:lang
	part     string  //路由中的一部分 例如 :lang
	children []*node //子节点 例如[doc, tutorial, intro]
	isWild   bool    //是否精确匹配 part 含有 : 或者 * true
}

//第一个匹配成功的节点，用于插入
func (n *node) matchChild(part string) *node {
	for _, child := range n.children {
		if child.part == part || child.isWild {
			return child
		}
	}

	return nil
}

//所有匹配成功的节点，用于查找
func (n *node) matchChildren(part string) []*node {
	nodes := make([]*node, 0)
	for _, child := range n.children {
		if child.part == part || child.isWild {
			nodes = append(nodes, child)
		}
	}

	return nodes
}

func (n *node) insert(pattern string, parts []string, height int) {
	if len(parts) == height {
		n.pattern = pattern
		return
	}

	part := parts[height]
	child := n.matchChild(part)
	if child == nil {
		child = &node{
			part:   part,
			isWild: part[0] == ':' || part[0] == '*',
		}
		n.children = append(n.children, child)
	}

	child.insert(pattern, parts, height+1)
}

func (n *node) search(parts []string, height int) *node {
	if len(parts) == height || strings.HasPrefix(n.part, "*") {
		if n.pattern == "" {
			return nil
		}

		return n
	}

	part := parts[height]
	children := n.matchChildren(part)

	for _, child := range children {
		res := child.search(parts, height+1) //查找匹配到的 part

		//spew.Dump(1111, res)
		//(*engine.node)(0xc0000583c0)({
		//pattern: (string) (len=12) "/hello/:name",
		//		part: (string) (len=5) ":name",
		//		children: ([]*engine.node) <nil>,
		//		isWild: (bool) true
		//})

		//(int) 1111
		//(*engine.node)(0xc000058440)({
		//pattern: (string) (len=17) "/assets/*filepath",
		//		part: (string) (len=9) "*filepath",
		//		children: ([]*engine.node) <nil>,
		//		isWild: (bool) true
		//})

		if res != nil {
			return res
		}
	}

	return nil
}
