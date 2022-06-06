package main

import "fmt"

//queue modified to hold address of Nodes
//binary tree is implemented with the
//help of Queue.

//++++++++++stack for tree traversals++++++++++++++++
type Stack struct {
	size int
	top  int
	arr  *[]*Node //will create the slice dyamically with size
}

func (st *Stack) createStack(size int) {
	//fmt.Println("Enter Size")
	//var size int      //store in size variable first , later assing to stack.size
	//fmt.Scanln(&size) //setting stack size
	st.size = size
	//fmt.Println("size = ", size)
	st.top = -1
	newSlice := make([]*Node, st.size)
	st.arr = &newSlice
}

func (st *Stack) DisplayStack() {
	for i := st.top; i >= 0; i-- {
		fmt.Printf("%d\n", (*st.arr)[i].data)
	}
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
	n := &Node{}
	if st.top == -1 {
		fmt.Println("Stack is empty")
	} else {
		n = (*st.arr)[st.top] //collect the deleting element
		st.top--              // reduce the top
	}
	return n
}

func (st *Stack) peek(index int) *Node { //looks in a index and return the element of that index
	n := &Node{}
	if st.top-index+1 < 0 {
		fmt.Println("invalid index")
	} else {
		n = (*st.arr)[st.top-index+1]
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

func (st *Stack) stackTop() *Node {
	if !st.isEmpty() {
		return (*st.arr)[st.top]
	}
	return nil
}

//----------stack for tree traversals----------------
//++++++++++tree using queue++++++
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

//now writing the iterative traversals functions

func preorderIterative(n *Node) {
	stack := &Stack{}
	stack.createStack(50)
	for n != nil || !stack.isEmpty() {
		if n != nil {
			fmt.Println(n.data)
			stack.push(n)
			n = n.left
		} else {
			n = stack.pop()
			n = n.right
		}
	}

}

func inorderIterative(n *Node) {
	stack := &Stack{}
	stack.createStack(50)
	for n != nil || !stack.isEmpty() {
		if n != nil {

			stack.push(n)
			n = n.left
		} else {
			n = stack.pop()
			fmt.Println(n.data)
			n = n.right
		}
	}

}

//------------tree using queue-------------
func main() {

	createTree()
	// fmt.Println("pre order display")
	// preorderDisplay(root)
	// fmt.Println("in order display")
	// inorderDisplay(root)
	// fmt.Println("post order display")
	// postorderDisplay(root)
	//uncomment above function calls for recursive traversals

	fmt.Println("iteraive preorder")
	preorderIterative(root)
	fmt.Println("iteraive inorder")
	inorderIterative(root)

}
