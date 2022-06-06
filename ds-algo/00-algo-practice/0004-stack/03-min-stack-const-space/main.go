package main

import "fmt"

//time O(1) and Space O(1) for the getMin function.
// For Push() : take a minEle variable. if an element > minEle no opereation nrequired, simply insert the value
//if some element < minEle, then apply the formula
//y(the modified value to be inserted into stack) = 2 * x(the value taken for push) - minEle
//by doing this the valueof y will be always smaller than minEle

//For Pop(): if the popped value y < minEle, that means this value was encrypted and inserted. so, return the minEle value as pop() return and
//calculate the new_minEle = 2 * x(old minEle) - y(popped element from the stack)
//if popped element is > minEle, then simply pop() returns that element from stack

// y = 2x - min (push)
//y + min = 2x
//min = 2x -y (pop)

type Stack struct {
	size        int
	topPosition int
	minEle      int
	arr         []int
}

var stack *Stack = &Stack{
	size:        10,
	topPosition: -1,
	minEle:      -1,
	arr:         make([]int, 10),
}

func (st *Stack) push(n int) {
	//check if stack is full
	if st.isFull() {
		fmt.Println("stack is full")
	} else {
		//prepare the very first insert
		if st.isEmpty() {
			st.topPosition = 0
			st.arr[st.topPosition] = n
			st.minEle = n
		} else {
			//insert until stack is full
			st.topPosition++
			toBeIserted := n
			if toBeIserted < st.minEle {
				toBeIserted = (2 * n) - st.minEle //as per above approach and formula
				st.minEle = n
				fmt.Println("after encryption of ", n, " tobe isnerted value = ", toBeIserted)
			}
			st.arr[st.topPosition] = toBeIserted
		}

	}

}

func (st *Stack) pop() int {
	x := -1
	//check is empty
	if st.isEmpty() {
		fmt.Println("stack is empty")
	} else {
		toBePopped := st.arr[st.topPosition]
		if toBePopped < st.minEle {
			x = st.minEle
			st.minEle = (2 * st.minEle) - (toBePopped)
		} else {

			x = toBePopped
		}
		st.arr[st.topPosition] = 0
		st.topPosition--
	}

	return x
}

func (st *Stack) isFull() bool {
	if st.topPosition == st.size-1 {
		return true
	}
	return false
}
func (st *Stack) isEmpty() bool {
	if st.topPosition == -1 {
		return true
	}
	return false
}

func (st *Stack) getMin() int {
	return 0
}

func main() {
	stack.push(18)
	fmt.Println("top = ", stack.topPosition, " minEle = ", stack.minEle, " array = ", stack.arr)
	stack.push(19)
	fmt.Println("top = ", stack.topPosition, " minEle = ", stack.minEle, " array = ", stack.arr)
	stack.push(12)
	fmt.Println("top = ", stack.topPosition, " minEle = ", stack.minEle, " array = ", stack.arr)
	stack.push(29)
	fmt.Println("top = ", stack.topPosition, " minEle = ", stack.minEle, " array = ", stack.arr)
	stack.push(10)
	fmt.Println("top = ", stack.topPosition, " minEle = ", stack.minEle, " array = ", stack.arr)
	stack.push(16)
	fmt.Println("top = ", stack.topPosition, " minEle = ", stack.minEle, " array = ", stack.arr)
	fmt.Println(stack.pop())
	fmt.Println("top = ", stack.topPosition, " minEle = ", stack.minEle, " array = ", stack.arr)
	fmt.Println(stack.pop())
	fmt.Println("top = ", stack.topPosition, " minEle = ", stack.minEle, " array = ", stack.arr)
	fmt.Println(stack.pop())
	fmt.Println("top = ", stack.topPosition, " minEle = ", stack.minEle, " array = ", stack.arr)
	fmt.Println(stack.pop())
	fmt.Println("top = ", stack.topPosition, " minEle = ", stack.minEle, " array = ", stack.arr)
}
