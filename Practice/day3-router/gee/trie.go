package gee

import (
	"fmt"
	"strings"
)

type node struct {
	// 待匹配路由
	pattern string
	// 路由中的一部分
	part string
	// 子节点
	children []*node
	// 是否精确匹配
	isWild bool
}

// matchChild 第一个匹配成功的节点 用于插入
func (n *node) matchChild(part string) *node {
	for _, child := range n.children {
		if child.part == part || child.isWild {
			return child
		}
	}
	return nil
}

// matchChildren 所有匹配成功的节点 用于搜索
func (n *node) matchChildren(part string) []*node {
	nodes := make([]*node, 0)
	for _, child := range n.children {
		if child.part == part || child.isWild {
			// func append(slice []Type, elems ...Type) []Type
			nodes = append(nodes, child)
		}
	}
	return nodes
}

// insert 前缀树节点插入
func (n *node) insert(pattern string, parts []string, height int) {
	// Array: the number of elements in v
	if len(parts) == height {
		n.pattern = pattern
		return
	}

	part := parts[height]
	child := n.matchChild(part)
	if child == nil {
		child = &node{part: part, isWild: part[0] == ':' || part[0] == '*'}
		n.children = append(n.children, child)
	}
	// 递归插入节点
	child.insert(pattern, parts, height+1)
}

// search 查询节点
func (n *node) search(parts []string, height int) *node {
	// strings.HasPrefix => return len(s) >= len(prefix) && s[0:len(prefix)] == prefix
	if len(parts) == height || strings.HasPrefix(n.part, "*") {
		if n.pattern == "" {
			return nil
		}
		return n
	}

	part := parts[height]
	children := n.matchChildren(part)

	for _, child := range children {
		result := child.search(parts, height+1)
		if result != nil {
			return result
		}
	}
	return nil
}

// String 打印节点信息
func (n *node) String() string {
	return fmt.Sprintf("node{pattern=%s, part=%s, isWild=%t}", n.pattern, n.part, n.isWild)
}

func (n *node) travel(list *([]*node)) {
	if n.pattern != "" {
		*list = append(*list, n)
	}
	for _, child := range n.children {
		child.travel(list)
	}
}
