package main

import (
	"fmt"
)

//chaining is a open hasing
type Node struct {
	Data int
	Next *Node
}

//this is single linked list
type LinkedList struct {
	Head   *Node
	Length int
}

var hashTable = make([]*Node, 10)

//has function
func Hash(key int) int {
	return key % 10
}

func SortedInsert(key int) {
	index := Hash(key)
	newNode := &Node{Data: key, Next: nil}
	if hashTable[index] == nil {
		//fmt.Println("index empty, directly insert into it")
		hashTable[index] = newNode
	} else {
		//some element is already present in the index
		//search for sorted insert position , then insert
		//considering no duplicate items
		//fmt.Println("node already present")
		newNode := &Node{Data: key, Next: nil}
		var positionNode *Node
		var onNode = hashTable[index]
		for onNode != nil && key > onNode.Data {
			positionNode = onNode
			onNode = onNode.Next
		}
		//fmt.Println("position node =", positionNode)
		//insert after the positionNode
		//check if the insert position is first
		if positionNode == nil {
			//fmt.Println("inserting in first position")
			newNode.Next = hashTable[index]
			hashTable[index] = newNode
		} else {
			//fmt.Println("inserting in calculated position")
			newNode.Next = positionNode.Next
			positionNode.Next = newNode
		}
	}
}

func PrintHashTable() {
	for i, v := range hashTable {
		node := v
		fmt.Printf("\nFor index %d:\n", i)
		for node != nil {
			fmt.Printf("\t%d", node.Data)
			node = node.Next
		}

	}
}

func Search(key int) bool {
	index := Hash(key)
	node := hashTable[index]
	flag := false
	for node != nil {
		if key == node.Data {
			flag = true
			break
		}
		node = node.Next
	}
	return flag
}

func Delete(key int) {
	index := Hash(key)
	headNode := hashTable[index]
	var positionNode *Node
	if headNode == nil { //nil check
		fmt.Println("empty index")
	} else if headNode.Data == key { //check the first node
		fmt.Println("found in first node..deleting")
		hashTable[index] = headNode.Next
	} else {
		for headNode.Next != nil { //check the rest of the sorted nodes

			if headNode.Next.Data == key {
				fmt.Println("deleting element ", headNode.Next.Data)
				positionNode = headNode //we've to delete the next node of positionNode
				break
			}
			headNode = headNode.Next
		}
		if positionNode != nil {
			positionNode.Next = positionNode.Next.Next
		} else {
			fmt.Println("element not found")
		}
	}
}

func main() {
	arr := []int{12, 22, 2, 32, 52, 42, 23, 63, 3, 5, 105, 95, 45, 65, 14, 4, 54, 64, 84, 94, 8}
	for _, v := range arr {
		SortedInsert(v)
	}
	PrintHashTable()

	//fmt.Println(Search(12))
	//Delete(12)
	//PrintHashTable()
	fmt.Println(Search(12))
	Delete(112)
	PrintHashTable()
}
