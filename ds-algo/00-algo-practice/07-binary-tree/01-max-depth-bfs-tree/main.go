package main

import (
	"fmt"
	"io"
	"os"
)

type Queue struct {
	size  int
	front int
	rear  int
	arr   []*Node
}

func (q *Queue) enqueue(n *Node) {
	if q.rear == q.size {
		fmt.Println("Queue is full")
	} else {
		q.rear++
		q.arr[q.rear] = n
	}
}

func (q *Queue) dequeue() *Node {
	var x *Node
	if q.front == q.rear {
		fmt.Println("Queue is empty")
		return nil
	} else {
		q.front++
		x = q.arr[q.front]
	}
	return x
}

func (q *Queue) isEmpty() bool {
	return q.front == q.rear
}

var queue = &Queue{
	size:  20,
	front: -1,
	rear:  -1,
	arr:   make([]*Node, 20),
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

type Node struct {
	left  *Node
	data  int
	right *Node
}

var root *Node

func insertRecursive(p *Node, key int) *Node {
	if p == nil {
		fmt.Println("creating a node for ", key)
		return &Node{data: key}
	}
	if key < p.data {
		p.left = insertRecursive(p.left, key)
	} else if key > p.data {
		p.right = insertRecursive(p.right, key)
	} else {
		return p
	}

	return p
}

func maxDepthRecursive(node *Node) int {
	if node == nil {
		return 0
	}
	x := maxDepthRecursive(node.left)
	y := maxDepthRecursive(node.right)

	if x > y {
		return x + 1
	} else {
		return y + 1
	}

}

func maxDepthBFSIterative(entryNode *Node) int {
	//use a queue for this function
	//breath first search approach
	if entryNode == nil {
		return 0
	}
	level := 0
	queue.enqueue(entryNode)

	for !queue.isEmpty() {
		levelNodes := dequeueAndPopulateNodes(queue)
		fmt.Println("returned level ", level, " nodes ", levelNodes)
		for _, node := range levelNodes {
			fmt.Println(node)
			if node.left != nil {
				queue.enqueue(node.left)
			}
			if node.right != nil {
				queue.enqueue(node.right)
			}

		}

		level++
	}
	fmt.Println("final level ", level)

	return level
}
func dequeueAndPopulateNodes(q *Queue) []*Node {
	var arr []*Node
	for !q.isEmpty() {
		arr = append(arr, q.dequeue())
	}
	return arr
}

func print(w io.Writer, node *Node, ns int, ch rune) {
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
	root = insertRecursive(root, 30)
	insertRecursive(root, 20)
	insertRecursive(root, 10)
	insertRecursive(root, 40)
	insertRecursive(root, 35)
	insertRecursive(root, 50)
	insertRecursive(root, 25)
	insertRecursive(root, 55)
	print(os.Stdout, root, 0, 'M')
	//fmt.Println(maxDepthRecursive(root))
	maxDepthBFSIterative(root)
}
