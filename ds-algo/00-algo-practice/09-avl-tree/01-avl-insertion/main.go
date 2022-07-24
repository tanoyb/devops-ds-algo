package main

import (
	"fmt"
	"io"
	"os"
)

//avl tree is a height balanced binary tree
type Node struct {
	left            *Node
	data            int
	heightOfSubtree int
	right           *Node
}

var root *Node = nil

func InsertRecursive(node *Node, key int) *Node {
	if node == nil {
		temp := &Node{data: key, heightOfSubtree: 1}
		return temp
	}
	if key < node.data {
		node.left = InsertRecursive(node.left, key)
	} else if key > node.data {
		node.right = InsertRecursive(node.right, key)
	}
	node.heightOfSubtree = GenNodeHeight(node)
	//now check the balance factor here
	//rotation is always checked on the new inserted node, in which side it has been added, the left or on the right side
	if BalanceFactor(node) == 2 && BalanceFactor(node.left) == 1 {
		//imbalance on left and then left, perform LL rotation
		return LL_Rotation(node)

	} else if BalanceFactor(node) == 2 && BalanceFactor(node.left) == -1 {
		//imbalance on left and then right, perform LR rotation
		return LR_Rotation(node)
	} else if BalanceFactor(node) == -2 && BalanceFactor(node.right) == -1 {
		//imbalance on right and then right, perform RR roation
	} else if BalanceFactor(node) == -2 && BalanceFactor(node.right) == 1 {
		//imbalance on right and then left, perform RL rotation
	}

	return node
}

func LR_Rotation(node *Node) *Node {
	nodeL := node.left
	nodeLtoR := nodeL.right

	//see the LR diagram to understand the assignment operations
	nodeL.right = nodeLtoR.left
	node.left = nodeLtoR.right
	nodeLtoR.left = nodeL
	nodeLtoR.right = node
	//now regenerate the changes in height
	nodeL.heightOfSubtree = GenNodeHeight(nodeL)
	node.heightOfSubtree = GenNodeHeight(node)
	nodeLtoR.heightOfSubtree = GenNodeHeight(nodeLtoR)
	if root == node {
		root = nodeLtoR
	}

	return nodeLtoR
}

func LL_Rotation(node *Node) *Node {
	nodeL := node.left
	nodeLtoR := nodeL.right
	nodeL.right = node
	node.left = nodeLtoR
	node.heightOfSubtree = GenNodeHeight(node)
	nodeL.heightOfSubtree = GenNodeHeight(nodeL)
	//now check if the node was the rootm if so then change the root to nodeL
	if root == node {
		root = nodeL
	}
	return nodeL
}

func BalanceFactor(node *Node) int {
	hleft, hright := 0, 0
	if node != nil && node.left != nil {
		hleft = node.left.heightOfSubtree
	}
	if node != nil && node.right != nil {
		hright = node.right.heightOfSubtree
	}
	return hleft - hright

}

func GenNodeHeight(node *Node) int {
	hleft := 0
	hright := 0
	if node != nil && node.left != nil {
		hleft = GenNodeHeight(node.left)
	}
	if node != nil && node.right != nil {
		hright = GenNodeHeight(node.right)
	}
	if hleft > hright {
		return hleft + 1
	} else {
		return hright + 1
	}
}

func PrintTree(w io.Writer, node *Node, space int, str string) {
	if node != nil {
		for i := 0; i < space; i++ {
			fmt.Fprintf(w, " ")
		}
		fmt.Fprintf(w, "%v:%v(%v)\n", str, node.data, node.heightOfSubtree)
		PrintTree(w, node.left, space+2, "Left")
		PrintTree(w, node.right, space+2, "Rght")
	}
}

func main() {
	// root = InsertRecursive(root, 10)
	// InsertRecursive(root, 5)
	// InsertRecursive(root, 2)
	root = InsertRecursive(root, 50)
	InsertRecursive(root, 10)
	InsertRecursive(root, 20)

	PrintTree(os.Stdout, root, 0, "Root")
}
