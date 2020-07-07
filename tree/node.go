package main

import "fmt"

type treeNode struct {
	value       int
	left, right *treeNode
}

func createNode(value int) *treeNode{
	return &treeNode{value:value}
}

func (node treeNode)print(){
	fmt.Print(node.value, " ")
}
func (node *treeNode) setValue(val int){
	if node == nil{
		fmt.Println("setting value to nil node")
		return
	}
	node.value = val
}
func (tree *treeNode)traverse(){
	if tree == nil{
		return
	}
	tree.left.traverse()
	fmt.Print(tree.value, " ")
	tree.right.traverse()
}
func main() {
	var root treeNode
	root = treeNode{value: 3}
	root.left = &treeNode{}
	root.right = &treeNode{5, nil, nil}
	root.right.left = new(treeNode)
	root.left.right = createNode((2))

	root.right.left.setValue(4)
	root.right.left.print()

	root.print()
	root.setValue(100)

	var pRoot *treeNode
	pRoot.setValue(200)
	pRoot = &root
	pRoot.setValue(300)
	pRoot.print()

	pRoot.traverse()
}
