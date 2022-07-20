package main

import (
	"fmt"
	"io"
	"os"
)

//https://www.youtube.com/watch?v=hTM3phVI6YQ&lc=UgxUajmaDDsmZ5Tp0U54AaABAg start from 5:46 using iterative method
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

func maxDepth(node *Node) int {
	hleft := 0
	hright := 0

	if node != nil && node.left != nil {
		hleft = maxDepth(node.left)
	}

	if node != nil && node.right != nil {
		hright = maxDepth(node.right)
	}
	if hleft > hright {
		return hleft + 1
	} else {
		return hright + 1
	}
}

func maxDepth1(node *Node) int {
	if node == nil {
		return 0
	}
	left := maxDepth(node.left)
	right := maxDepth(node.right)
	if left > right {
		return left + 1
	} else {

		return right + 1
	}
}

func main() {
	root = InsertRecursive(root, 30)
	InsertRecursive(root, 20)
	InsertRecursive(root, 10)
	InsertRecursive(root, 40)
	InsertRecursive(root, 35)
	InsertRecursive(root, 50)
	InsertRecursive(root, 25)
	InsertRecursive(root, 60)

	myPrintTreeRecursive(os.Stdout, root, 0, "Root")
	fmt.Println("==max depth===", maxDepth(root))
	fmt.Println("==max depth===", maxDepth1(root))

}
