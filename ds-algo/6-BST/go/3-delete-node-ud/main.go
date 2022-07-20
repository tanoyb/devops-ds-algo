package main

import "fmt"

//when deleting from a BST,
//first search for the element, if found then delete
//if the found node is a leaf node, then delete is easy
//if the found node has  one child, move that child to the
//parent of that deleted node
//if the found node has both left and right children
//then, the inorder predecessor/inorder successor with take take the place
//of that deleted node. see that diagram for more clear explanation

type Node struct {
	leftChild  *Node
	data       int
	rightChild *Node
}

var root *Node = nil

func insertNode(key int) { //iterative insert
	t := root
	r := &Node{}
	p := &Node{}
	if root == nil {
		//create root node
		p.data = key
		root = p
		return
	}
	for t != nil {
		r = t
		if key < t.data {
			t = t.leftChild
		} else if key > t.data {
			t = t.rightChild
		} else {
			return
		}
	}
	newNode := &Node{data: key}
	if key < r.data {
		r.leftChild = newNode
	} else {
		r.rightChild = newNode
	}
}

func InsertRecursive(node *Node, key int) *Node { //this is correct
	if node == nil {
		temp := &Node{data: key}
		return temp
	}
	if key < node.data {
		node.leftChild = InsertRecursive(node.leftChild, key)
	} else if key > node.data {
		node.rightChild = InsertRecursive(node.rightChild, key)
	}

	return node
}

func InorderTraversal(node *Node) {
	if node != nil {
		InorderTraversal(node.leftChild)
		fmt.Println(node.data)
		InorderTraversal(node.rightChild)
	}
}

func Search(key int) bool {
	t := root
	for t != nil {
		if key == t.data {
			return true
		} else if key < t.data {
			t = t.leftChild
		} else {
			t = t.rightChild
		}
	}
	return false
}

func HeightOfTree(node *Node) int {
	if node == nil {
		return 0
	}
	x := HeightOfTree(node.leftChild)
	y := HeightOfTree(node.rightChild)

	if x > y {
		return x + 1
	} else {
		return y + 1
	}
}

func InorderPredecessor(node *Node) *Node {

	if node == nil {
		return nil
	}
	for node != nil && node.rightChild != nil {
		node = node.rightChild
	}

	return node
}

func InorderSuccessor(node *Node) *Node {
	if node == nil {
		return nil
	}
	for node != nil && node.leftChild != nil {
		node = node.leftChild
	}

	return node
}

func DeleteNode(p *Node, key int) *Node {
	if p == nil {
		return nil
	}
	if p.leftChild == nil && p.rightChild == nil {
		if p == root {
			root = nil
		}
		p = nil
		return nil
	}

	if key < p.data {
		p.leftChild = DeleteNode(p.leftChild, key)
	} else if key > p.data {
		p.rightChild = DeleteNode(p.rightChild, key)
	} else {
		//element found. now delete based on the scenarios
		if HeightOfTree(p.leftChild) > HeightOfTree(p.rightChild) {
			//find inorder predecessor
			q := InorderPredecessor(p.leftChild)
			p.data = q.data
			p.leftChild = DeleteNode(p.leftChild, q.data)
		} else {
			//find inorder successor
			q := InorderSuccessor(p.rightChild)
			p.data = q.data
			p.rightChild = DeleteNode(p.rightChild, q.data)
		}
	}
	return p

}

func main() {
	root = InsertRecursive(root, 10)
	InsertRecursive(root, 5)
	InsertRecursive(root, 20)
	InsertRecursive(root, 8)
	InsertRecursive(root, 30)

	InorderTraversal(root)
	fmt.Println()
	DeleteNode(root, 20)
	fmt.Println()
	InorderTraversal(root)
	fmt.Println()
	fmt.Println(Search(30))
	fmt.Println(Search(40))
}
