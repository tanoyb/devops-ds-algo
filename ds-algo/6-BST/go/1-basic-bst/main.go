package main

import "fmt"

//Node
type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

var count int

//Insert
func (n *Node) Insert(k int) {
	if n.Key < k {
		//move right
		if n.Right == nil {
			n.Right = &Node{Key: k}
		} else {
			n.Right.Insert(k)
		}
	} else if n.Key > k {
		//move left
		if n.Left == nil {
			n.Left = &Node{Key: k}
		} else {
			n.Left.Insert(k)
		}
	}
}

func (n *Node) Search(k int) bool {
	count++
	if n == nil {
		return false
	} else if n.Key < k {
		//move right
		return n.Right.Search(k)

	} else if n.Key > k {
		//move left
		return n.Left.Search(k)
	}
	return true //that means n.Key == k
}

func main() {
	tree := &Node{Key: 100}
	fmt.Println(tree)
	tree.Insert(50)
	tree.Insert(110)
	tree.Insert(120)
	tree.Insert(99)
	tree.Insert(90)
	tree.Insert(27)
	tree.Insert(45)
	fmt.Println(tree)
	fmt.Println(tree.Search(45))
	fmt.Println(count)
}
