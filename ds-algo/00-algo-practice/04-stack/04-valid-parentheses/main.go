package main

import (
	"fmt"
	"strings"
)

//stack valid parentheses
//============= stack code starts from here============
type Node struct {
	data string
	next *Node
}

type StackLinkedList struct {
	top_or_head *Node
	size        int
}

var top *Node = &Node{next: nil}

func push(n string) {
	tempNode := &Node{data: n}
	if (Node{} == *tempNode) {
		fmt.Println("Stack/heap is full")
	} else {
		tempNode.next = top
		top = tempNode
	}

}

func (s *StackLinkedList) Size() int {
	return s.size
}

func (s *StackLinkedList) IsEmpty() bool {
	return s.size == 0
}

func (s *StackLinkedList) Peek() string {
	if s.IsEmpty() {
		fmt.Println("empty stack")
		return ""
	}
	return s.top_or_head.data
}

func (s *StackLinkedList) Push(n string) {
	tempNode := &Node{data: n}
	if (Node{} == *tempNode) {
		fmt.Println("Stack/heap is full")
	} else {
		tempNode.next = s.top_or_head
		s.top_or_head = tempNode
		s.size++
	}
	// s.top_or_head = &Node{data: n, next: s.top_or_head}
	// s.size++
}

func (s *StackLinkedList) Pop() (string, bool) {
	if s.IsEmpty() {
		fmt.Println("stack is empty")
		return "", false
	}
	value := s.top_or_head.data
	s.top_or_head = s.top_or_head.next
	s.size--
	return value, true
}

func (s *StackLinkedList) Display() {
	temp := s.top_or_head
	for temp != nil {
		fmt.Printf("%s\n", temp.data)
		temp = temp.next
	}
	fmt.Println()
}

//============= stack code ends here============
//actual problem code starts
var stack StackLinkedList
var parenthesesMap map[string]string = map[string]string{
	")": "(",
	"}": "{",
	"]": "[",
}

func checkValidParentheses(s string) bool {
	str := strings.Split(s, "")
	strOpen := "({["
	strClose := ")}]"
	//fmt.Println(str)
	for i, v := range str {
		fmt.Println("i=", i, "=v==", v)
		if strings.Contains(strClose, v) && i == 0 {
			fmt.Println("starting with close brackets")
			return false
		} else {
			if strings.Contains(strClose, v) {
				//received a closing breacket
				//check if the stack is empty
				if stack.IsEmpty() {
					return false
				} else {
					popped, _ := stack.Pop()
					if parenthesesMap[v] != popped {
						return false
					}
				}
			}
		}
		if strings.Contains(strOpen, v) {
			stack.Push(v)
		}

	}
	if !stack.IsEmpty() {
		return false
	}

	return true
}

func main() {
	s := "((({}{}[(({}))])())){}[]"
	fmt.Println(checkValidParentheses(s))
}
