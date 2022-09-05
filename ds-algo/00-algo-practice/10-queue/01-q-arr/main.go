package main

import "fmt"

//queue using an array and two pointers
type Queue struct {
	front int
	rear  int
	size  int
	arr   []int
}

func (q *Queue) Enqueue(n int) {
	//check if queue is full
	if q.front == q.size {
		fmt.Println("queue is full")
	} else {
		q.rear += 1
		q.arr[q.rear] = n
		fmt.Println("element pushed")
	}
}

func (q *Queue) DeQueue() int {
	//check if queue is empty
	x := -1
	if q.front == q.rear {
		fmt.Println("Queue is empty")
	} else {
		q.front += 1
		x = q.arr[q.front]
		fmt.Println("element de-queued")
	}
	return x
}

func main() {
	queue := &Queue{
		front: -1,
		rear:  -1,
		size:  5,
		arr:   make([]int, 5),
	}

}
