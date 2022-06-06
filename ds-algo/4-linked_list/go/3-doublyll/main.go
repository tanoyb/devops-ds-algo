package main

import (
	"fmt"
	"os"
	"strconv"
)

type Node struct {
	Data int
	Prev *Node
	Next *Node
}

type doublyLL struct {
	Head   *Node
	Tail   *Node
	Length int
}

func (dll *doublyLL) InsertAtEnd(num int) {
	tempNode := &Node{Data: num, Next: nil, Prev: nil}
	if dll.Head == nil && dll.Tail == nil {
		dll.Head = tempNode
		dll.Tail = tempNode
	} else {
		tempNode.Prev = dll.Tail
		dll.Tail.Next = tempNode
		dll.Tail = tempNode
	}

}

func (dll doublyLL) Display() {
	var p *Node = dll.Head
	for p != nil {
		fmt.Printf("%d \n", p.Data)
		p = p.Next
	}
}

func (dll doublyLL) DisplayReverse() {
	var p *Node = dll.Tail
	for p != nil {
		fmt.Printf("%d \n", p.Data)
		p = p.Prev
	}
}

func (dll *doublyLL) InsertAtBeginning(num int) {
	tempNode := &Node{Data: num, Prev: nil, Next: nil}
	if dll.Head == nil && dll.Tail == nil {
		dll.Head = tempNode
		dll.Tail = tempNode
	} else {
		tempNode.Next = dll.Head
		dll.Head.Prev = tempNode
		dll.Head = tempNode
	}
}

func (dll *doublyLL) DeleteFromBeginning() {
	if dll.Head == dll.Tail {
		dll.Head = nil
		dll.Tail = nil
	} else {
		dll.Head = dll.Head.Next
		dll.Head.Prev = nil
	}
	dll.Length--
}

func (dll *doublyLL) DeleteFromEnd() {
	if dll.Head == dll.Tail {
		dll.Head = nil
		dll.Tail = nil
	} else {
		dll.Tail = dll.Tail.Prev
		dll.Tail.Next = nil
	}
	dll.Length--
}

func main() {
	dll := doublyLL{}
	var choice string
	for true {
		fmt.Println("Enter your choice")
		fmt.Println("1. Insert node at the end")
		fmt.Println("2. traverse/display")
		fmt.Println("3. traverse/display in reverse order")
		fmt.Println("4. insert at beginning")
		fmt.Println("5. delete from beginning")
		fmt.Println("6. delete from end")
		fmt.Println("7. delete by value(not available)")
		fmt.Println("8. exit")
		fmt.Scanln(&choice)
		switch choice {
		case "1":
			var data string
			fmt.Println("enter number to insert")
			fmt.Scanln(&data)
			num, _ := strconv.Atoi(data)
			dll.InsertAtEnd(num)

		case "2":
			dll.Display()
		case "3":
			dll.DisplayReverse()
		case "4":
			var data string
			fmt.Println("enter number to insert at beginning")
			fmt.Scanln(&data)
			num, _ := strconv.Atoi(data)
			dll.InsertAtBeginning(num)
		case "5":
			dll.DeleteFromBeginning()
		case "6":
			dll.DeleteFromEnd()
		default:
			os.Exit(0)
		}
	}

}
