package main

import "fmt"

type CircularQueue struct {
	Size  int
	Front int
	Rear  int
	Arr   []int
}

func (q *CircularQueue) Enqueue(n int) {
	fmt.Println(" enqueue ", n)
	fmt.Println("front before=", q.Front, " rear before=", q.Rear)
	fmt.Println("array before = ", q.Arr)

	if (q.Rear+1)%q.Size == q.Front {
		fmt.Println("Queue is full")
	} else {
		q.Rear = (q.Rear + 1) % q.Size
		q.Arr[q.Rear] = n
	}
	fmt.Println("front after=", q.Front, " rear after=", q.Rear)
	fmt.Println("array after = ", q.Arr)
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
	fmt.Println("front after=", q.Front, " rear after=", q.Rear)
	fmt.Println("array after = ", q.Arr)
	return x
}

func main() {
	queueArr := &CircularQueue{
		Size:  5,
		Front: 0,
		Rear:  0,
		Arr:   make([]int, 5),
	}
	//notice that we are pointing the front to 0 and we are not using the 0 index in case of circular queue
	//fmt.Println(queueArr.Arr)
	queueArr.Enqueue(10)
	queueArr.Enqueue(20)
	queueArr.Enqueue(30)
	queueArr.Enqueue(40)
	queueArr.Enqueue(50)
	fmt.Println(queueArr.Dequeue())

	queueArr.Enqueue(60)
	queueArr.Enqueue(70)
	// queueArr.Enqueue(6)
	fmt.Println(queueArr.Dequeue())
	fmt.Println(queueArr.Dequeue())
	fmt.Println(queueArr.Dequeue())
	fmt.Println(queueArr.Dequeue())
	fmt.Println(queueArr.Dequeue())
}
