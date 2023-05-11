package main

import "fmt"

//circular queue
//move front and rear accordingly.
//wherever the front is pointing must be an empty space
//because the rear and front will be equal and it will point that queue is empty.
//front==rear -> empty condition
//if size of the arr is 10, we will store for 9 places
// 1 space is left for the front pointer. because front should be left empty
//use mod-remainder (by arr size) operation for obtaining the circular index values
//because when rear comes to the end of the arr, it will
//again go to the begging of the arr to insert values.
//next place of rear == front => queue is full

type CircularQueue struct {
	size  int
	front int
	rear  int
	arr   []int
}

func (q *CircularQueue) enqueue(n int) {
	if (q.rear+1)%q.size == q.front {
		fmt.Println("Queue is full")
	} else {
		q.rear = (q.rear + 1) % q.size
		q.arr[q.rear] = n
	}
}

func (q *CircularQueue) dequeue() int {
	x := -1
	if q.front == q.rear {
		fmt.Println("Queue is empty")
		return -1
	} else {
		q.front = (q.front + 1) % q.size
		x = q.arr[q.front]
	}
	return x
}

func (q *CircularQueue) display() { //needs to be modified
	for i := q.front + 1; i <= q.rear; i++ {
		fmt.Println(q.arr[i])
	}
}

func main() {
	queue := &CircularQueue{}
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

/*
package main

import "fmt"

type Queue struct {
	size  int
	front int
	rear  int
	arr   *[]int
}

func (q *Queue) enqueue(n int) {
	if (q.rear+1)%q.size == q.front {
		fmt.Println("queue is full")
	} else {
		q.rear = (q.rear + 1) % q.size
		(*q.arr)[q.rear] = n
		fmt.Println(n, "inserted")
	}
}

func (q *Queue) dequeue() {
	if q.front == q.rear {
		fmt.Println("queue is empty")
	} else {
		q.front = (q.front + 1) % q.size
		x := (*q.arr)[q.front]
		(*q.arr)[q.front] = 0
		fmt.Println(x, " removed from queue")

	}
}

func main() {
	a := make([]int, 7)
	q := &Queue{
		size:  7,
		front: 0,
		rear:  0,
		arr:   &a,
	}
	q.enqueue(1)
	q.enqueue(2)
	q.enqueue(3)
	q.enqueue(4)
	q.enqueue(5)
	q.enqueue(6)
	q.enqueue(7)
	fmt.Println("arr = ", (*q.arr), "  front == ", q.front, " rear == ", q.rear)
	q.dequeue()
	q.dequeue()
	q.dequeue()
	q.dequeue()
	q.enqueue(7)
	q.enqueue(8)
	q.dequeue()
	q.dequeue()
	q.dequeue()
	q.dequeue()
	q.dequeue()

	fmt.Println("arr = ", (*q.arr), "  front == ", q.front, " rear == ", q.rear)

}

*/
