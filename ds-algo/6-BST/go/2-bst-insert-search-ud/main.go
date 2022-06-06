package main

import "fmt"

type Node struct {
	leftChild  *Node
	data       int
	rightChild *Node
}

var root *Node = &Node{}

func insertNode(key int) {
	t := root
	r := &Node{}
	p := &Node{}
	if root == nil {
		//create root node
		p.data = key
		root = p
		return
	}
	for t != nil {
		r = t
		if key < t.data {
			t = t.leftChild
		} else if key > t.data {
			t = t.rightChild
		} else {
			return
		}
	}
	newNode := &Node{data: key}
	if key < r.data {
		r.leftChild = newNode
	} else {
		r.rightChild = newNode
	}
}

func InorderTraversal(node *Node) {
	if node != nil {
		InorderTraversal(node.leftChild)
		fmt.Println(node.data)
		InorderTraversal(node.rightChild)
	}
}

func Search(key int) bool {
	t := root
	for t != nil {
		if key == t.data {
			return true
		} else if key < t.data {
			t = t.leftChild
		} else {
			t = t.rightChild
		}
	}
	return false
}

func main() {
	insertNode(10)
	insertNode(5)
	insertNode(20)
	insertNode(8)
	insertNode(30)

	InorderTraversal(root)
	fmt.Println(Search(30))
	fmt.Println(Search(40))
}
