package main

import (
	"fmt"

	"golang.org/x/tools/container/intsets"

	"github.com/gostudy/tree"

)

// go get golang.org/x/tools/cmd/goimports
// go get -v github.com/gpmgo/gopm
//组合
type myTreeNode struct {
	node *tree.Node
}

func (myNode *myTreeNode)  postOrder (){

	if myNode == nil || myNode.node == nil {
		return
	}


	left := myTreeNode{myNode.node.Left}
	left.postOrder()
	right := myTreeNode{myNode.node.Right}
	right.postOrder()
	myNode.node.Print()
}

func testSparse()  {
	s := intsets.Sparse{}
	s.Insert(1)
	s.Insert(1000)
	s.Insert(1000000)
	fmt.Println(s.Has(1000))
	fmt.Println(s.Has(10000))
}


func main() {

	var root tree.Node

	root = tree.Node{Value: 3}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{Value: 5}
	root.Right.Left = new(tree.Node)

	root.Left.Right = tree.CreateNode(2)

	root.Right.Left.SetValue(4)

	//实现练习
	root.Traverse()

	nodeCount := 0
	root.TraverSeFunc(func(node *tree.Node) {
		nodeCount++
	})

	fmt.Println("Node count: ", nodeCount)

	//fmt.Println()
	//myRoot := myTreeNode{&root}
	//myRoot.postOrder()
	//fmt.Println()
	//
	//testSparse()




}

