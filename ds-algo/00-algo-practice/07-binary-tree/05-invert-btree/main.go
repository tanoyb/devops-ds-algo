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

func invertBTree(node *Node) *Node {
	if node == nil {
		return nil
	}

	temp := node.left
	node.left = node.right
	node.right = temp
	invertBTree(node.left)
	invertBTree(node.right)
	return node
}

func main() {
	root = InsertRecursive(root, 30)
	InsertRecursive(root, 20)
	InsertRecursive(root, 10)
	InsertRecursive(root, 40)
	InsertRecursive(root, 35)
	InsertRecursive(root, 50)
	InsertRecursive(root, 25)

	myPrintTreeRecursive(os.Stdout, root, 0, "Root")
	invertBTree(root)
	myPrintTreeRecursive(os.Stdout, root, 0, "Root")
}
