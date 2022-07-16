package main

import (
	"fmt"
	"io"
	"os"
)

type Node struct {
	left  *Node
	data  int
	right *Node
}

var root *Node = nil

func insertIterative(n int) {
	p := root
	//check if root is nil
	if root == nil {
		node := &Node{data: n}
		root = node
		return
	}
	var last *Node = nil
	for p != nil {
		last = p
		if n > p.data {
			//move to right
			p = p.right
		} else if n < p.data {
			//move to left
			p = p.left
		} else {
			return
		}
	}
	//now we have the insert location node last
	newNode := &Node{data: n}
	if n < last.data {
		last.left = newNode
	} else {
		last.right = newNode
	}
}

func InsertRecursive(node *Node, key int) *Node {
	if node == nil {
		temp := &Node{data: key}
		return temp
	}
	if key < node.data {
		node.left = InsertRecursive(node.left, key)
	} else if key > node.data {
		node.right = InsertRecursive(node.right, key)
	}

	return node
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

func main() {
	// insertIterative(50)
	// insertIterative(40)
	// insertIterative(60)
	// insertIterative(45)
	// insertIterative(55)
	// insertIterative(40)

	root = InsertRecursive(root, 30)
	InsertRecursive(root, 20)
	InsertRecursive(root, 10)
	InsertRecursive(root, 40)
	InsertRecursive(root, 35)
	InsertRecursive(root, 50)
	InsertRecursive(root, 25)

	myPrintTreeRecursive(os.Stdout, root, 0, "Root")
}
