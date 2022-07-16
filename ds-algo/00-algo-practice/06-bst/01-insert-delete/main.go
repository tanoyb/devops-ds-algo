package main

import (
	"fmt"
)

type Node struct {
	left  *Node
	data  int
	right *Node
}

var root *Node

func insert(key int) {
	fmt.Println("trying to insert ", key)
	newNode := &Node{data: key}
	var positionNode *Node
	currentNode := root
	if root == nil {
		fmt.Println("inserted into root node")
		root = newNode
	} else {
		for currentNode != nil {
			positionNode = currentNode
			fmt.Println("on current node ", currentNode)
			if key == currentNode.data {
				//return
				fmt.Println("element alrady present in the tree")
			} else if key < currentNode.data {
				//go left
				fmt.Println("going to left")
				currentNode = currentNode.left
			} else {
				//right
				fmt.Println("going to right")
				currentNode = currentNode.right
			}
		}
		fmt.Println("position node is ", positionNode)
		if key < positionNode.data {
			positionNode.left = newNode
		} else {
			positionNode.right = newNode
		}
	}
}

func insertRecursive(p *Node, key int) *Node {
	if p == nil {
		fmt.Println("creating a node for ", key)
		return &Node{data: key}
	}
	if key < p.data {
		p.left = insertRecursive(p.left, key)
	} else if key > p.data {
		p.right = insertRecursive(p.right, key)
	}

	return p
}

func deleteRecursive(p *Node, key int) *Node {
	fmt.Println("trying to delete ", key)
	//here p is the root node.
	if p == nil {
		return nil
	}
	if p.left == nil && p.right == nil {
		if p == root {
			root = nil
		}
		return nil
	}
	var temp *Node
	if key < p.data {
		p.left = deleteRecursive(p.left, key)
	} else if key > p.data {
		p.right = deleteRecursive(p.right, key)
	} else {
		//element has been found
		//check height of inorder predessor or successor
		if heightOfSubtree(p.left) >= heightOfSubtree(p.right) {
			//with inorder predessor
			temp = InPre(p.left)
			p.data = temp.data
			p.left = deleteRecursive(p.left, temp.data)
		} else {
			temp = InSucc(p.right)
			p.data = temp.data
			p.right = deleteRecursive(p.right, temp.data)
		}
	}
	return p
}

func heightOfSubtree(p *Node) int {

	if p == nil {
		return 0
	}
	x := heightOfSubtree(p.left)
	y := heightOfSubtree(p.right)
	if x > y {
		return x + 1
	} else {
		return y + 1
	}

}

func InPre(p *Node) *Node {
	if p == nil {
		return nil
	}

	for p != nil && p.right != nil {
		p = p.right
	}

	return p
}

func InSucc(p *Node) *Node {
	if p == nil {
		return nil
	}
	for p != nil && p.left != nil {
		p = p.left
	}
	return p
}

func main() {
	// insert(30)
	// insert(20)
	// insert(10)
	// insert(40)
	// insert(35)
	// insert(50)
	// insert(4)
	root = insertRecursive(root, 30)
	insertRecursive(root, 20)
	insertRecursive(root, 10)
	insertRecursive(root, 40)
	insertRecursive(root, 35)
	insertRecursive(root, 50)
	insertRecursive(root, 25)

	//deleteRecursive(root, 30)
}
