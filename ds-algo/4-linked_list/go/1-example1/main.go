package main

import "fmt"

type node struct {
	data int
	next *node
}

type LinkedList struct {
	head   *node
	length int
}

func (l *LinkedList) prepend(n *node) {
	second := l.head
	l.head = n
	l.head.next = second
	l.length++
}

func (l LinkedList) printListData() {
	head := l.head
	for l.length != 0 {
		fmt.Printf("%d \n", head.data)
		head = head.next
		l.length--
	}
}

func (l *LinkedList) deleteNodeWithValue(v int) {
	previousHeadNode := l.head
	for previousHeadNode.next.data != v {
		previousHeadNode = previousHeadNode.next
	}
	//https://www.youtube.com/watch?v=8QoynPUY9_8
}

func main() {
	newlist := LinkedList{}
	node1 := &node{data: 10}
	node2 := &node{data: 20}
	node3 := &node{data: 30}
	newlist.prepend(node1)
	newlist.prepend(node2)
	newlist.prepend(node3)
	fmt.Println(newlist)
	newlist.printListData()
}
