package main

import "fmt"

//==========stack implementation ++++++++++++========
type Stack struct {
	size int
	top  int
	arr  []int //will create the slice dyamically with size
}

func (st *Stack) DisplayStack() {
	for i := st.top; i >= 0; i-- {
		fmt.Printf("%d\n", st.arr[i])
	}
}

func (st *Stack) push(n int) {
	if st.top == st.size-1 {
		fmt.Println("Stack is full")
	} else {
		st.top++
		st.arr[st.top] = n
	}
}

func (st *Stack) pop() int {
	n := -1
	if st.top == -1 {
		fmt.Println("Stack is empty")
	} else {
		n = st.arr[st.top] //collect the deleting element
		st.top--           // reduce the top
	}
	return n
}

func (st *Stack) peek(index int) int { //looks in a index and return the element of that index
	n := -1
	if st.top-index+1 < 0 {
		fmt.Println("invalid index")
	} else {
		n = st.arr[st.top-index+1]
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

func (st *Stack) stackTop() int {
	if !st.isEmpty() {
		return st.arr[st.top]
	}
	return -1
}

////==========stack implementation ------------========

var originalStack *Stack = &Stack{
	size: 10,
	top:  -1,
	arr:  make([]int, 10),
}

var minElementStack *Stack = &Stack{
	size: 10,
	top:  -1,
	arr:  make([]int, 10),
}

func insertInStack(n int) {
	if originalStack.isFull() {
		fmt.Println("stack is full")
	} else if originalStack.isEmpty() {
		originalStack.push(n)
		minElementStack.push(n)
	} else {
		originalStack.push(n)
		minTop := minElementStack.stackTop()
		if n < minTop {
			minElementStack.push(n)
		} else {
			minElementStack.push(minTop)
		}
	}
	//fmt.Println(originalStack)
	//fmt.Println(minElementStack)
}

func getMin() int {

	return minElementStack.stackTop()
}

func main() {
	insertInStack(9)
	fmt.Println(getMin())
	insertInStack(5)
	fmt.Println(getMin())
	insertInStack(12)
	fmt.Println(getMin())
	insertInStack(3)
	fmt.Println(getMin())
	insertInStack(17)
	fmt.Println(getMin())
	insertInStack(1)
	fmt.Println(getMin())
	insertInStack(10)
	fmt.Println(getMin())
	insertInStack(11)
	fmt.Println(getMin())
	insertInStack(12)
	fmt.Println(getMin())
	insertInStack(0)
	fmt.Println(getMin())

	//now adjust the pop function to pop from both the stacks
}
