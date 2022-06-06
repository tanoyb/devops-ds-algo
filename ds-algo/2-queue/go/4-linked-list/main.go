package main

import (
	"fmt"
)

//using front and rear pointer

type Node struct {
	data int
	next *Node
}

type Queue struct {
	front *Node
	rear  *Node
	size  int
}

func (q *Queue) enqueue(n int) {
	node := new(Node) //using new to create new node
	//node := &Node{}
	//but new is not suggested, we can use struct litereal to create a node in a better way
	//fmt.Print(node)

	if node == nil {
		fmt.Println("Memory full...queue is full")
	} else {
		node.data = n
		node.next = nil
		if q.front == nil {
			q.front = node
			q.rear = node
		} else {
			q.rear.next = node
			q.rear = node

		}
	}
}

func (q *Queue) dequeue() int {
	x := -1
	if q.front == nil {
		fmt.Println("Queue is empty")
	} else {
		x = q.front.data
		q.front = q.front.next
	}

	return x
}

func (q *Queue) display() {
	//node := &Node{}
	node := q.front
	for node != nil {
		fmt.Println(node.data)
		node = node.next
	}
}

func main() {
	queue := &Queue{}
	//fmt.Println(queue)
	queue.enqueue(10)
	queue.enqueue(20)
	queue.enqueue(30)
	queue.display()
	queue.dequeue()
	queue.display()

}
