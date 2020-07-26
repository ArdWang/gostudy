package tree

import (
	"fmt"
)

type Node struct {
	Value int
	Left, Right *Node
}

func (node Node) Print() {
	fmt.Println(node.Value," ")
}

func (node *Node) SetValue(value int) {
	if node == nil {
		fmt.Println("Setting value to nil node. Ignored")
		return
	}
	node.Value = value
}

func (node *Node) TraverseFunc(f func(*Node)) {

}




func CreateNode(value int) *Node {
	return &Node{Value: value}
}

//func main() {
//	var root treeNode
//
//	root = treeNode{value: 3}
//	root.left = &treeNode{}
//	root.right = &treeNode{5,nil,nil}
//	root.right.left = new(treeNode)
//
//	root.left.right = createNode(2)
//
//	//nodes :=[] treeNode{
//	//	{value: 3},
//	//	{},
//	//	{6,nil,&root},
//	//}
//	//
//	//fmt.Println(nodes)
//
//	//root.print()
//	//fmt.Println()
//
//	root.right.left.setValue(4)
//
//	//实现练习
//	root.traverse()
//
//
//	//root.right.left.print()
//	//fmt.Println()
//	//
//	//root.print()
//	//root.setValue(100)
//	//
//	////pRoot := &root
//	////pRoot.print()
//	////pRoot.setValue(200)
//	////pRoot.print()
//	//
//	//var pRoot *treeNode
//	//pRoot.setValue(200)
//	//pRoot = &root
//	//pRoot.setValue(300)
//	//pRoot.print()
//
//}
