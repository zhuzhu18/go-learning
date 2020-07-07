package main

import (
	"fmt"
	tree "tree3"
)

type myTreeNode struct {
	node *tree.TreeNode
}

func (myNode *myTreeNode) postOrder() {
	if myNode == nil || myNode.node == nil {
		return
	}
	var left = myTreeNode{myNode.node.Left}
	left.postOrder()
	var right = myTreeNode{myNode.node.Right}
	right.postOrder()
	myNode.node.Print()
}

func main() {
	var root tree.TreeNode
	root = tree.TreeNode{Value: 3}
	root.Left = &tree.TreeNode{}
	root.Right = &tree.TreeNode{5, nil, nil}
	root.Right.Left = new(tree.TreeNode)
	root.Left.Right = tree.CreateNode(2)

	root.Right.Left.SetValue(4)

	var pRoot *tree.TreeNode
	pRoot = &root
	pRoot.Print()
	fmt.Println()
	pRoot.Traverse()

	fmt.Println()
	myRoot := myTreeNode{&root}
	myRoot.postOrder()
}
