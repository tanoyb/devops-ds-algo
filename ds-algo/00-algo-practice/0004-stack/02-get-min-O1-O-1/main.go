package main

import "fmt"

//implements getMIn operation in contant time without
//using the auxliary array.

//approach : we use the same array, with even and odd index
//use the even indexes(0,2,4,6..) to store the value and use
//odd indexes (1,3,5,7..) to store the coresponding min element
//see the diagram for more clear idea

type Stack struct {
	size                  int
	topPosition           int
	currentMinTopPosition int
	arr                   []int //will create the slice dyamically with size
}

func (st *Stack) push(n int) {
	fmt.Println("trying to push ", n)
	if st.isFull() {
		fmt.Println("Stack is full")
	} else if st.isEmpty() { //very first insert
		fmt.Println("very first entry")
		st.topPosition = 0
		st.arr[st.topPosition] = n
		//now put the current min on the one index above top
		st.currentMinTopPosition = st.topPosition + 1
		st.arr[st.currentMinTopPosition] = n

	} else { //for insertion from second time onwards....

		st.topPosition = st.topPosition + 2
		st.arr[st.topPosition] = n
		//now put the current min on the one index above top
		minValue := st.arr[st.currentMinTopPosition]
		if n < minValue {
			st.currentMinTopPosition = st.topPosition + 1
			st.arr[st.currentMinTopPosition] = n
		} else {
			st.currentMinTopPosition = st.topPosition + 1
			st.arr[st.currentMinTopPosition] = minValue
		}

	}

}

func (st *Stack) pop() int {

	n := -1
	if st.topPosition == -1 {
		fmt.Println("Stack is empty")
	} else if st.topPosition == 0 { //check if the element is last element
		n = st.arr[st.topPosition]
		st.arr[st.topPosition] = 0
		st.arr[st.currentMinTopPosition] = 0
		st.topPosition = -1
		st.currentMinTopPosition = -1

	} else {
		n = st.arr[st.topPosition] //collect the deleting element
		st.arr[st.topPosition] = 0
		st.arr[st.currentMinTopPosition] = 0
		st.topPosition = st.topPosition - 2 // reduce the top
		st.currentMinTopPosition = st.currentMinTopPosition - 2
	}
	fmt.Println("trying to pop ", n)
	return n
}

func (st *Stack) peek(index int) int { //looks in a index and return the element of that index
	n := -1
	if st.topPosition-index+1 < 0 {
		fmt.Println("invalid index")
	} else {
		n = st.arr[st.topPosition-index+1]
	}
	return n

}

func (st *Stack) isEmpty() bool {
	if st.topPosition == -1 {
		return true
	}
	return false
}

func (st *Stack) isFull() bool {
	if st.topPosition == st.size-2 {
		return true
	}
	return false
}

func (st *Stack) stackTop() int {
	if !st.isEmpty() {
		return st.arr[st.topPosition]
	}
	return -1
}

var stack *Stack = &Stack{
	size:                  22, // index 0,1,3.21 = 22.  taken even number because the last min value at index 21 with hold the min value for index 20
	topPosition:           -1,
	currentMinTopPosition: -1,
	arr:                   make([]int, 22),
}

func (st *Stack) getMin() int {
	return st.arr[st.currentMinTopPosition]
}
func main() {
	stack.push(9)
	fmt.Println("top = ", stack.topPosition, " minTop = ", stack.currentMinTopPosition, " array = ", stack.arr)
	stack.push(10)
	fmt.Println("top = ", stack.topPosition, " minTop = ", stack.currentMinTopPosition, " array = ", stack.arr)
	stack.push(7)
	fmt.Println("top = ", stack.topPosition, " minTop = ", stack.currentMinTopPosition, " array = ", stack.arr)
	stack.push(12)
	fmt.Println("top = ", stack.topPosition, " minTop = ", stack.currentMinTopPosition, " array = ", stack.arr)
	stack.push(1)
	fmt.Println("top = ", stack.topPosition, " minTop = ", stack.currentMinTopPosition, " array = ", stack.arr)
	fmt.Println(stack.getMin())
	stack.pop()
	fmt.Println("top = ", stack.topPosition, " minTop = ", stack.currentMinTopPosition, " array = ", stack.arr)
	fmt.Println(stack.getMin())
	stack.pop()
	fmt.Println("top = ", stack.topPosition, " minTop = ", stack.currentMinTopPosition, " array = ", stack.arr)
	fmt.Println(stack.getMin())
	stack.pop()
	fmt.Println("top = ", stack.topPosition, " minTop = ", stack.currentMinTopPosition, " array = ", stack.arr)
	fmt.Println(stack.getMin())

}
