package tree

import "fmt"

type TreeNode struct {
	Value       int
	Left, Right *TreeNode
}

func CreateNode(value int) *TreeNode {
	return &TreeNode{Value: value}
}

func (node TreeNode) Print() {
	fmt.Print(node.Value, " ")
}
func (node *TreeNode) SetValue(val int) {
	if node == nil {
		fmt.Println("setting value to nil node")
		return
	}
	node.Value = val
}
func (tree *TreeNode) Traverse() {
	if tree == nil {
		return
	}
	tree.Left.Traverse()
	fmt.Print(tree.Value, " ")
	tree.Right.Traverse()
}