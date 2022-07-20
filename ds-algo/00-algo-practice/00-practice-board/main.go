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

func heightOfSubtree(p *Node) int {
	if p == nil {
		return 0
	}
	heightLeft := heightOfSubtree(p.left)
	heightRight := heightOfSubtree(p.right)
	if heightLeft > heightRight {
		return heightLeft + 1
	} else {
		return heightRight + 1
	}
}

func inorderPredecessor(node *Node) *Node {

	if node == nil {
		return nil
	}
	for node != nil && node.right != nil {
		node = node.right
	}

	return node
}

func inorderSuccessor(node *Node) *Node {
	if node == nil {
		return nil
	}
	for node != nil && node.left != nil {
		node = node.left
	}

	return node
}

func deleteRecursive(p *Node, key int) *Node {
	if p == nil {
		return nil
	}
	if p.left == nil && p.right == nil {
		if p == root {
			root = nil
		}
		p = nil
		return nil
	}

	if key < p.data {
		p.left = deleteRecursive(p.left, key)
	} else if key > p.data {
		p.right = deleteRecursive(p.right, key)
	} else {
		//the element has found
		if heightOfSubtree(p.left) > heightOfSubtree(p.right) {
			//find inorder predecessor
			inPre := inorderPredecessor(p.left)
			p.data = inPre.data
			p.left = deleteRecursive(p.left, inPre.data)
		} else {
			inSucc := inorderSuccessor(p.right)
			p.data = inSucc.data
			p.right = deleteRecursive(p.right, inSucc.data)
		}

	}

	return p
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
