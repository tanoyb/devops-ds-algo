package main

import "fmt"

//queue modified to hold address of Nodes
//binary tree is implemented with the
//help of Queue.
type Node struct {
	left  *Node
	data  int
	right *Node
}

type Queue struct {
	size  int
	front int
	rear  int
	arr   []*Node
}

var root *Node = &Node{}
var queue *Queue = &Queue{}

func (q *Queue) enqueue(n *Node) {
	if q.rear == q.size {
		fmt.Println("Queue is full")
	} else {
		q.rear++
		q.arr[q.rear] = n
	}
}

func (q *Queue) dequeue() *Node {
	x := &Node{}
	if q.front == q.rear {
		fmt.Println("Queue is empty")
		return x
	} else {
		q.front++
		x = q.arr[q.front]
	}
	return x
}

func (q *Queue) display() {
	for i := q.front + 1; i <= q.rear; i++ {

		temp := q.arr[i]
		fmt.Println(*temp)
	}
}

func (q *Queue) isEmpty() bool {
	return q.front == q.rear
}
func initQueue(size int) {
	queue.size = size
	queue.front, queue.rear = -1, -1
	queue.arr = make([]*Node, size)
}

func createTree() {

	x := -1
	initQueue(50)
	fmt.Println("Enter root integer value")
	fmt.Scanf("%d", &x)
	root = &Node{data: x, left: nil, right: nil}
	queue.enqueue(root)

	for !queue.isEmpty() {
		dequeuedNode := queue.dequeue()
		//for left child
		fmt.Println("Enter left child (integer) value of ", dequeuedNode.data)
		fmt.Scanf("%d", &x)
		if x != -1 {
			tempLeftNode := &Node{data: x, left: nil, right: nil}
			dequeuedNode.left = tempLeftNode
			queue.enqueue(tempLeftNode)
		}

		//for right child
		fmt.Println("Enter right child (integer) value of ", dequeuedNode.data)
		fmt.Scanf("%d", &x)
		if x != -1 {
			tempRightNode := &Node{data: x, left: nil, right: nil}
			dequeuedNode.right = tempRightNode
			queue.enqueue(tempRightNode)
		}

	}

}

func preorderDisplay(n *Node) {
	if n != nil {
		fmt.Println(n.data)
		preorderDisplay(n.left)
		preorderDisplay(n.right)
	}
}

func inorderDisplay(n *Node) {
	if n != nil {
		inorderDisplay(n.left)
		fmt.Println(n.data)
		inorderDisplay(n.right)
	}
}

func postorderDisplay(n *Node) {
	if n != nil {
		postorderDisplay(n.left)
		postorderDisplay(n.right)
		fmt.Println(n.data)
	}
}

func main() {
	// initQueue(50)
	// queue.enqueue(&Node{data: 10})
	// queue.enqueue(&Node{data: 20})
	// queue.enqueue(&Node{data: 30})
	// queue.enqueue(&Node{data: 40})
	// queue.enqueue(&Node{data: 50})
	// fmt.Println()
	// queue.display()
	// queue.dequeue()
	// fmt.Println()
	// queue.display()
	// queue.dequeue()
	// fmt.Println()
	// queue.display()
	// fmt.Println(queue.isEmpty())
	createTree()
	fmt.Println("pre order display")
	preorderDisplay(root)
	fmt.Println("in order display")
	inorderDisplay(root)
	fmt.Println("post order display")
	postorderDisplay(root)

}
