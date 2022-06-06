package queue

import "fmt"

type CircularQueue struct {
	Size  int
	Front int
	Rear  int
	Arr   []int
}

func (q *CircularQueue) Enqueue(n int) {
	//fmt.Println(" enqueue ", n)
	//fmt.Println("front before=", q.Front, " rear before=", q.Rear)
	//fmt.Println("array before = ", q.Arr)

	if (q.Rear+1)%q.Size == q.Front {
		fmt.Println("Queue is full")
	} else {
		q.Rear = (q.Rear + 1) % q.Size
		q.Arr[q.Rear] = n
	}
	//fmt.Println("front after=", q.Front, " rear after=", q.Rear)
	//fmt.Println("array after = ", q.Arr)
}

func (q *CircularQueue) Dequeue() int {
	x := -1
	if q.Front == q.Rear {
		fmt.Println("Queue is empty")
		return -1
	} else {
		q.Front = (q.Front + 1) % q.Size
		x = q.Arr[q.Front]
	}
	//fmt.Println("front after=", q.Front, " rear after=", q.Rear)
	//fmt.Println("array after = ", q.Arr)
	return x
}

func (q *CircularQueue) IsEmpty() bool {
	return q.Front == q.Rear
}
