package main

import (
	"fmt"
	"learngo/tree"
)

type myTreeNode struct {
	node *tree.Node
}

func (myNode *myTreeNode) postOrder() {
	if myNode == nil || myNode.node == nil {
		return
	}

	left := myTreeNode{myNode.node.Left}
	left.postOrder()
	right := myTreeNode{myNode.node.Right}
	right.postOrder()
	myNode.node.Print()
}

func main()  {
	var root tree.Node
	root = tree.Node{Value: 3}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5, nil, nil}
	//root.right.left = new(Node)
	root.Right.Left = tree.CreateTreeNode(2)

	nodes := []tree.Node{
		{Value: 3},
		{},
		{6, nil, &root},
	}

	fmt.Println(root)
	fmt.Println(nodes)

	root.Print()
	fmt.Print("\n")

	root.Right.SetValue(100)
	root.Right.Print()

	root.Right.Left.SetValue(-20)
	root.Right.Left.Print()

	pRoot := &root
	pRoot.SetValue(300)
	pRoot.Print()
	//
	node := myTreeNode{&root}
	node.postOrder()
	fmt.Println()
}