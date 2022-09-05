package main

import "fmt"

type CircularQueue struct {
	Size  int
	Front int
	Rear  int
	Arr   []int
}

func (q *CircularQueue) Enqueue(n int) {
	//check if queue is full
	if (q.Rear+1)%q.Size == q.Front {
		fmt.Println("queue is full")
	} else {
		q.Rear = (q.Rear + 1) % q.Size
		q.Arr[q.Rear] = n
		fmt.Println("element pushed")
	}
}

func (q *CircularQueue) Dequeue() int {
	x := -1
	//check if queue is empty
	if q.Front == q.Rear {
		fmt.Println("Queue is empty")
		//once queue is empty, resetting the front and rear pointer to initial position
		q.Front = 0
		q.Rear = 0
	} else {
		q.Front = (q.Front + 1) % q.Size
		x = q.Arr[q.Front]
	}
	return x
}

func main() {

	cirQueue := &CircularQueue{
		Size:  5,
		Front: 0,
		Rear:  0,
		Arr:   make([]int, 5),
	}
	cirQueue.Enqueue(10)
}
