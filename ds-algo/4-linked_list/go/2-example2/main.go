package main

import "fmt"

type Node struct {
	Data int
	Next *Node
}

//this is single linked list
type LinkedList struct {
	Head   *Node
	Length int
}

//taken pointer receiver because you will work on the actual linked list
func (l *LinkedList) prepend(n *Node) {
	second := l.Head
	l.Head = n
	l.Head.Next = second
	l.Length++
}

func (l *LinkedList) append(n *Node) {

}

func (l *LinkedList) deleteWithValue(i int) int {
	if l.Length == 0 {
		fmt.Println("Empty list")
		return 0
	}
	fmt.Printf("trying to delete %d \n", i)
	prevTempNode := l.Head
	if l.Head.Data == i {
		l.Head = l.Head.Next
		l.Length--
		return i
	}

	for prevTempNode.Next.Data != i {
		if prevTempNode.Next.Next == nil {
			return 0
		}
		prevTempNode = prevTempNode.Next
	}
	if prevTempNode.Next.Next == nil && prevTempNode.Next.Data == i {
		prevTempNode.Next = nil
		l.Length--
		return i
	}
	prevTempNode.Next = prevTempNode.Next.Next
	l.Length--
	return i

}

//taken value pointer because we don't need to modify anything
//on the linked list, we just need to print the values, that means
//we will work on the copy of the linked list.
func (l LinkedList) printListData() {
	toPrint := l.Head
	len := l.Length
	for len != 0 {
		fmt.Printf(" value = %d \n", toPrint.Data)
		toPrint = toPrint.Next
		len--
	}
}

func main() {
	mylist := LinkedList{}
	fmt.Println(mylist)
	node1 := &Node{Data: 20}
	node2 := &Node{Data: 30}
	node3 := &Node{Data: 40}
	node4 := &Node{Data: 50}
	node5 := &Node{Data: 60}
	node6 := &Node{Data: 70}
	mylist.prepend(node1)
	mylist.prepend(node2)
	mylist.prepend(node3)
	mylist.prepend(node4)
	mylist.prepend(node5)
	mylist.prepend(node6)
	fmt.Println(mylist)
	mylist.printListData()
	mylist.deleteWithValue(20)
	fmt.Printf("after delete==========\n")
	mylist.printListData()
}
