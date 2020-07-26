package tree

import "fmt"

func (node *Node) Traverse() {

	node.TraverSeFunc(func(n *Node) {
		n.Print()
	})
	fmt.Println()
}

func (node *Node) TraverSeFunc (f func(*Node)){
	if node == nil {
		return
	}

	node.Left.TraverseFunc(f)
	f(node)
	node.Right.TraverseFunc(f)
}
