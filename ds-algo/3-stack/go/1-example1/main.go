package main

import "fmt"

type Stack struct {
	items []int
}

//push adds a value at the end
func (s *Stack) Push(i int) {
	s.items = append(s.items, i)
}

//pop removes a value from the end
func (s *Stack) Pop() int {
	lastIndex := len(s.items) - 1
	toRemove := s.items[lastIndex]
	s.items = s.items[:lastIndex]
	return toRemove
}

func main() {
	myStack := Stack{}
	fmt.Println(myStack)
	myStack.Push(1)
	myStack.Push(2)
	myStack.Push(3)
	myStack.Push(4)
	fmt.Println(myStack)
	myStack.Pop()
	fmt.Println(myStack)
}
