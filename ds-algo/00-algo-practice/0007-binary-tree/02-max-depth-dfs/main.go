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

var root *Node

func insertRecursiveWithDepth(p *Node, key, level int) *Node {

	if p == nil {
		//fmt.Println("creating a node for ", key)
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

func insertIterativeWithDepth(key int) {
	//fmt.Println("trying to insert ", key)
	//newNode := &Node{data: key}
	var positionNode *Node
	currentNode := root
	level := 0
	if root == nil {
		//fmt.Println("inserted into root node")
		root = &Node{data: key, depth: level}
	} else {
		for currentNode != nil {
			positionNode = currentNode
			//fmt.Println("on current node ", currentNode)
			if key == currentNode.data {
				//return
				fmt.Println("element alrady present in the tree")
			} else if key < currentNode.data {
				//go left
				//fmt.Println("going to left")
				currentNode = currentNode.left
			} else {
				//right
				//fmt.Println("going to right")
				currentNode = currentNode.right
			}
			level = level + 1
		}
		//fmt.Println("position node is ", positionNode)
		if key < positionNode.data {
			positionNode.left = &Node{data: key, depth: level}
		} else {
			positionNode.right = &Node{data: key, depth: level}
		}
	}
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

var starr = make([]*Node, 10)
var stack = &Stack{
	size: 10,
	top:  -1,
	arr:  &starr,
}

func maxDepthDFSIterative(node *Node) int { //use preorder iterative for max depth
	//this approach uses a stack in epth first search approach
	//each node has its depth within it.
	//https://www.youtube.com/watch?v=hTM3phVI6YQ
	level := 0
	for node != nil || !stack.isEmpty() {
		if node != nil {
			if node.depth > level {
				level = node.depth
			}
			stack.push(node)
			node = node.left
		} else {
			lastNode := stack.pop()
			node = lastNode.right

		}

	}

	return level
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

func myPrintTreeRecursive(w io.Writer, node *Node, space int, child string) {
	if node != nil {
		for i := 0; i < space; i++ {
			fmt.Fprintf(w, " ")
		}
		fmt.Fprintf(w, "%v:%v-%v\n", child, node.data, node.depth)
		myPrintTreeRecursive(w, node.left, space+2, "Left")
		myPrintTreeRecursive(w, node.right, space+2, "Rigt")
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
	insertIterativeWithDepth(60)
	insertIterativeWithDepth(22)
	insertIterativeWithDepth(75)
	insertIterativeWithDepth(18)
	insertIterativeWithDepth(80)
	insertIterativeWithDepth(14)
	insertIterativeWithDepth(85)
	insertIterativeWithDepth(5)
	insertIterativeWithDepth(95)
	myPrintTreeRecursive(os.Stdout, root, 0, "Root")
	fmt.Println("max depth is == ", maxDepthDFSIterative(root))
}
