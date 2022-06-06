package main

import (
	"fmt"
	"io"
	"os"
)

type BinaryNode struct {
	left  *BinaryNode
	right *BinaryNode
	data  int
}

type BinaryTree struct {
	root *BinaryNode
}

func (t *BinaryTree) insert(n int) *BinaryTree {
	if t.root == nil { //first node in the tree
		t.root = &BinaryNode{data: n, left: nil, right: nil}
	} else {
		t.root.insertNode(n)
	}
	return t
}

func (n *BinaryNode) insertNode(x int) {
	if n == nil {
		return
	} else if x <= n.data {
		if n.left == nil {
			n.left = &BinaryNode{data: x, left: nil, right: nil}
		} else {
			n.left.insertNode(x)
		}
	} else {
		if n.right == nil {
			n.right = &BinaryNode{data: x, left: nil, right: nil}
		} else {
			n.right.insertNode(x)
		}
	}
}

func print(w io.Writer, node *BinaryNode, ns int, ch rune) {
	if node == nil {
		return
	}
	for i := 0; i < ns; i++ {
		fmt.Fprintf(w, " ")
	}
	fmt.Fprintf(w, "%c:%v\n", ch, node.data)
	print(w, node.left, ns+2, 'L')
	print(w, node.right, ns+2, 'R')
}

func main() {
	tree := &BinaryTree{}
	tree.insert(100).
		insert(-20).
		insert(-50).
		insert(-15).
		insert(-60).
		insert(50).
		insert(60).
		insert(55).
		insert(85).
		insert(15).
		insert(5).
		insert(-10)
	print(os.Stdout, tree.root, 0, 'M')
}
