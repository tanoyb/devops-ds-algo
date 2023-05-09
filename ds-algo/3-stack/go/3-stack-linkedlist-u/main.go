package main

import (
	"fmt"
)

type Node struct {
	data int
	next *Node
}

type StackLinkedList struct {
	top_or_head *Node
	size        int
}

func (s *StackLinkedList) Size() int {
	return s.size
}

func (s *StackLinkedList) IsEmpty() bool {
	return s.size == 0
}

func (s *StackLinkedList) Peek() int {
	if s.IsEmpty() {
		fmt.Println("empty stack")
		return 0
	}
	return s.top_or_head.data
}

func (s *StackLinkedList) Push(n int) {
	tempNode := &Node{data: n} //in case of heap is full, the node creation will fail and tempNode will be nil
	if (Node{} == *tempNode) { //check if tempNode is empty using Node{} literal
		fmt.Println("Stack/heap is full")
	} else {
		tempNode.next = s.top_or_head
		s.top_or_head = tempNode
		s.size++
	}
	// s.top_or_head = &Node{data: n, next: s.top_or_head}
	// s.size++
}

func (s *StackLinkedList) Pop() (int, bool) {
	if s.IsEmpty() {
		fmt.Println("stack is empty")
		return 0, false
	}
	value := s.top_or_head.data
	s.top_or_head = s.top_or_head.next
	s.size--
	return value, true
}

func (s *StackLinkedList) Display() {
	temp := s.top_or_head
	for temp != nil {
		fmt.Printf("%d\n", temp.data)
		temp = temp.next
	}
	fmt.Println()
}

func main() {
	var stack StackLinkedList
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	stack.Push(5)
	stack.Push(6)
	stack.Push(7)
	stack.Push(8)
	stack.Display()
	stack.Pop()
	stack.Pop()
	stack.Pop()
	stack.Pop()
	stack.Pop()
	stack.Pop()
	stack.Pop()
	stack.Pop()
	stack.Pop()
	stack.Display()
}

//go run -gcflags -m main.go -to see stack and heap allocation
