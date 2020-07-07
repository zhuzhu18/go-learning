package main

import (
	"fmt"
)

type TreeNode struct {
	Value       int
	Left, Right *TreeNode
}

func CreateNode(val int) *TreeNode {
	return &TreeNode{Value: val}
}

func (node *TreeNode) SetValue(val int) {
	if node == nil {
		fmt.Println("setting value to nil node")
		return
	}
	node.Value = val
}
func (node *TreeNode) TraverseFunc(f func(node *TreeNode)) {
	if node == nil {
		return
	}
	node.Left.TraverseFunc(f)
	f(node)
	node.Right.TraverseFunc(f)
}

func (node *TreeNode) TraverseWithChannel() chan *TreeNode {
	out := make(chan *TreeNode)
	go func() {
		node.TraverseFunc(func(node *TreeNode) {
			out <- node // 将每个节点的数据发送给channel
		})
		close(out) // 发送完后关闭该channel
	}()
	return out
}
func main() {
	var root TreeNode
	root = TreeNode{Value: 3}
	root.Left = &TreeNode{}
	root.Right = &TreeNode{5, nil, nil}
	root.Right.Left = new(TreeNode)
	root.Left.Right = CreateNode(2)
	root.Right.Left.SetValue(4)

	c := root.TraverseWithChannel()
	maxNode := 0
	for node := range c { // 从channel接收数据
		if node.Value > maxNode {
			maxNode = node.Value
		}
	}
	fmt.Println("Max node value:", maxNode)
}
