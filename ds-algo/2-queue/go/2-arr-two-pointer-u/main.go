package main

import "fmt"

//using front and rear pointer(index)
//front should point before the first element
//when front==rear, queue is empty
//rear pointer -> for inserting
//front pointer for deletion

type Queue struct {
	size  int
	front int
	rear  int
	arr   []int
}

func (q *Queue) enqueue(n int) {
	if q.rear == q.size {
		fmt.Println("Queue is full")
	} else {
		q.rear++
		q.arr[q.rear] = n
	}
}

func (q *Queue) dequeue() int {
	x := -1
	if q.front == q.rear {
		fmt.Println("Queue is empty")
		return -1
	} else {
		q.front++
		x = q.arr[q.front]
	}
	return x
}

func (q *Queue) display() {
	for i := q.front + 1; i <= q.rear; i++ {
		fmt.Println(q.arr[i])
	}
}

func main() {
	queue := &Queue{}
	queue.size = 5
	queue.front, queue.rear = -1, -1
	queue.arr = make([]int, queue.size)

	//creation complete , now call functions
	queue.enqueue(10)
	queue.enqueue(20)
	queue.enqueue(30)
	queue.enqueue(40)
	queue.enqueue(50)
	queue.display()
	queue.dequeue()
	queue.display()
	queue.dequeue()
	queue.dequeue()
	queue.dequeue()
	queue.dequeue()
	queue.dequeue()

}
