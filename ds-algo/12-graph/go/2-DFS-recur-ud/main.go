package main

import "fmt"

//==========stack implementation ++++++++++++========
type Node struct {
	Data int
	Next *Node
}

type StackLinkedList struct {
	Top_or_head *Node
	Size        int
}

func (s *StackLinkedList) Push(n int) {
	tempNode := &Node{Data: n}
	if (Node{} == *tempNode) {
		fmt.Println("Stack/heap is full")
	} else {
		tempNode.Next = s.Top_or_head
		s.Top_or_head = tempNode
		s.Size++
	}
	// s.top_or_head = &Node{data: n, next: s.top_or_head}
	// s.size++
}

func (s *StackLinkedList) Pop() int {
	if s.IsEmpty() {
		fmt.Println("stack is empty")
		return -1
	}
	value := s.Top_or_head.Data
	s.Top_or_head = s.Top_or_head.Next
	s.Size--
	return value
}

func (s *StackLinkedList) IsEmpty() bool {
	return s.Size == 0
}

////==========stack implementation ------------========
var node = &Node{}
var stack *StackLinkedList = &StackLinkedList{
	Top_or_head: node,
	Size:        0, //size does not matter for linke lists
}

var AdjMatrix [][]int = [][]int{
	{0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 1, 1, 1, 0, 0, 0},
	{0, 1, 0, 1, 0, 0, 0, 0},
	{0, 1, 1, 0, 1, 1, 0, 0},
	{0, 1, 0, 1, 0, 1, 0, 0},
	{0, 0, 0, 1, 1, 0, 1, 1},
	{0, 0, 0, 0, 0, 1, 0, 0},
	{0, 0, 0, 0, 0, 1, 0, 0},
}
var visited [10]int
var noOfVerteces = 7

func DFS_recursive(u int) {
	if visited[u] == 0 {
		fmt.Println(u)
		visited[u] = 1
		for v := 1; v <= noOfVerteces; v++ {
			if AdjMatrix[u][v] == 1 && visited[v] == 0 {
				DFS_recursive(v)
			}
		}
	}
}

func main() {
	DFS_recursive(1)

}
