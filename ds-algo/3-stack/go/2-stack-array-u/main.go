package main

import "fmt"

type Stack struct {
	size int
	top  int
	arr  *[]int //will create the slice dyamically with size
}

func (st *Stack) create() {
	fmt.Println("Enter Size")
	var size int      //store in size variable first , later assing to stack.size
	fmt.Scanln(&size) //setting stack size
	st.size = size
	//fmt.Println("size = ", size)
	st.top = -1
	newSlice := make([]int, st.size)
	st.arr = &newSlice
}

func (st *Stack) DisplayStack() {
	for i := st.top; i >= 0; i-- {
		fmt.Printf("%d\n", (*st.arr)[i])
	}
}

func (st *Stack) push(n int) {
	if st.top == st.size-1 {
		fmt.Println("Stack is full")
	} else {
		st.top++
		(*st.arr)[st.top] = n
	}
}

func (st *Stack) pop() int {
	n := -1
	if st.top == -1 {
		fmt.Println("Stack is empty")
	} else {
		n = (*st.arr)[st.top] //collect the deleting element
		st.top--              // reduce the top
	}
	return n
}

func (st *Stack) peek(index int) int { //looks in a index and return the element of that index
	n := -1
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

func (st *Stack) stackTop() int {
	if !st.isEmpty() {
		return (*st.arr)[st.top]
	}
	return -1
}

func main() {
	st := &Stack{}
	st.create()
	st.push(1)
	st.push(2)
	st.push(3)
	st.push(4)
	st.push(5)
	st.push(6)
	fmt.Println("stack size==", st.size)
	fmt.Println("stack top==", st.top)
	st.DisplayStack()
	fmt.Println("stack top==", st.top)
	st.pop()
	st.DisplayStack()
	fmt.Println("stack top==", st.top)
	st.pop()
	st.DisplayStack()
	fmt.Println("stack top==", st.top)
	st.pop()
	st.pop()
	st.pop()
	st.DisplayStack()
	fmt.Println("stack top==", st.top)
	fmt.Println(st.peek(1))
	fmt.Println(st.peek(30))

}
