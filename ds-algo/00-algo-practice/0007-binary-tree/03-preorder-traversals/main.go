package main

import (
	"fmt"
	"io"
	"os"
)

type Node struct {
	left  *Node
	data  int
	depth int
	right *Node
}

type Stack struct {
	size int
	top  int
	arr  *[]*Node //will create the slice dyamically with size
}

func (st *Stack) push(n *Node) {
	if st.top == st.size-1 {
		fmt.Println("Stack is full")
	} else {
		st.top++
		(*st.arr)[st.top] = n
	}
}

func (st *Stack) pop() *Node {
	var n *Node
	if st.top == -1 {
		fmt.Println("Stack is empty")
	} else {
		n = (*st.arr)[st.top] //collect the deleting element
		st.top--              // reduce the top
	}
	return n
}
func (st *Stack) isEmpty() bool {
	if st.top == -1 {
		return true
	}
	return false
}

func (st *Stack) isFull() bool {
	if st.top == st.size-1 {
		return true
	}
	return false
}

var starr = make([]*Node, 50)
var stack = &Stack{
	size: 50,
	top:  -1,
	arr:  &starr,
}

var root *Node

func insertRecursiveWithDepth(p *Node, key, level int) *Node {

	if p == nil {
		fmt.Println("creating a node for ", key)
		return &Node{data: key, depth: level}
	}

	if key < p.data {
		p.left = insertRecursiveWithDepth(p.left, key, level+1)
	} else if key > p.data {
		p.right = insertRecursiveWithDepth(p.right, key, level+1)
	} else {
		return p
	}

	return p
}

func myPrintTreeRecursive(w io.Writer, node *Node, space int, child string) {
	if node != nil {
		for i := 0; i < space; i++ {
			fmt.Fprintf(w, " ")
		}
		fmt.Fprintf(w, "%v:%v\n", child, node.data)
		myPrintTreeRecursive(w, node.left, space+2, "Left")
		myPrintTreeRecursive(w, node.right, space+2, "Rigt")
	}
}

func preorderRecursive(node *Node) {
	if node != nil {
		fmt.Println(node.data)
		preorderRecursive(node.left)
		preorderRecursive(node.right)
	}
}

func preorderIterative(node *Node) {
	//iterative version has to use a stack

	for node != nil || !stack.isEmpty() {
		if node != nil {
			fmt.Println(node.data)
			stack.push(node)
			node = node.left
		} else {
			lastNode := stack.pop()
			node = lastNode.right

		}

	}

}

func main() {
	root = insertRecursiveWithDepth(root, 30, 0)
	insertRecursiveWithDepth(root, 20, 0)
	insertRecursiveWithDepth(root, 10, 0)
	insertRecursiveWithDepth(root, 40, 0)
	insertRecursiveWithDepth(root, 35, 0)
	insertRecursiveWithDepth(root, 50, 0)
	insertRecursiveWithDepth(root, 25, 0)
	insertRecursiveWithDepth(root, 55, 0)
	insertRecursiveWithDepth(root, 60, 0)
	insertRecursiveWithDepth(root, 22, 0)
	insertRecursiveWithDepth(root, 75, 0)
	insertRecursiveWithDepth(root, 18, 0)
	insertRecursiveWithDepth(root, 80, 0)
	insertRecursiveWithDepth(root, 14, 0)
	insertRecursiveWithDepth(root, 85, 0)
	insertRecursiveWithDepth(root, 5, 0)
	insertRecursiveWithDepth(root, 57, 0)

	myPrintTreeRecursive(os.Stdout, root, 0, "Root")
	fmt.Println("============")
	preorderRecursive(root)
	fmt.Println("============")
	preorderIterative(root)

}
