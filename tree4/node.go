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
	tree.TraverseFunc(func(n *TreeNode){
		n.Print()
	})
	fmt.Println()
}

func (node *TreeNode)TraverseFunc(f func(*TreeNode)){
	if node == nil {
		return
	}
	node.Left.TraverseFunc(f)
	f(node)
	node.Right.TraverseFunc(f)
}