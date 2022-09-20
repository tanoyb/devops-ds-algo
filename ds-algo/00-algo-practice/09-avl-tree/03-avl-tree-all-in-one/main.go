package main

import (
	"fmt"
	"io"
	"os"
)

//AVL tree, add node, delete node, rotations recursive and iterative

type Node struct {
	left          *Node
	right         *Node
	subtreeHeight int
	data          int
}

var root *Node = nil

func InsertIteratvie(key int) bool {
	//fmt.Println("trying to insert ", key)
	newNode := &Node{data: key}
	currentNode := root
	if root == nil {
		root = newNode
	} else {
		//find the inserting position first
		var position *Node

		for currentNode != nil {

			//fmt.Println("current node ", currentNode.data)
			position = currentNode
			if key == currentNode.data {
				return false
			} else if key < currentNode.data {
				//go left

				currentNode = currentNode.left
			} else if key > currentNode.data {
				//go right

				currentNode = currentNode.right
			}

		}
		//fmt.Println("The position node is ", position.data)
		if key < position.data {
			position.left = newNode
		} else {
			position.right = newNode
		}
		//fmt.Printf("node %v inserted\n\n", key)
	}
	return true
}

func InsertRecursive(node *Node, key int) *Node {
	if node == nil {
		temp := &Node{data: key}
		return temp
	}
	if key < node.data {
		node.left = InsertRecursive(node.left, key)
	} else if key > node.data {
		node.right = InsertRecursive(node.right, key)
	}
	node.subtreeHeight = NodeHeight(node)
	fmt.Println("BF node = ", BalanceFactor(node), " and BF node.left=", BalanceFactor(node.left), " and BF node.right=", BalanceFactor(node.right))
	if BalanceFactor(node) == 2 && BalanceFactor(node.left) == 1 {
		//LL rotation, L then L imbalance.
		return LL_Rotation(node)

	} else if BalanceFactor(node) == 2 && BalanceFactor(node.left) == -1 {
		//LR rrotation, L then R imbalance
		return LR_Rotation(node)

	} else if BalanceFactor(node) == -2 && BalanceFactor(node.right) == -1 {
		//RR rotation, R then R imbalance
		return RR_Rotation(node)
	} else if BalanceFactor(node) == -2 && BalanceFactor(node.right) == 1 {
		//RL rotation, R then L imbalance
		return RL_Rotation(node)
	}
	return node
}

func RL_Rotation(node *Node) *Node {
	fmt.Println("performing RL rotation")
	// nodeLeft := node.left
	// nodeLeftRight := nodeLeft.right
	// nodeLeft.right = nodeLeftRight.left
	// node.left = nodeLeftRight.right
	// nodeLeftRight.left = nodeLeft
	// nodeLeftRight.right = node
	nodeRight := node.right
	nodeRightLeft := nodeRight.left

	nodeRight.left = nodeRightLeft.right
	node.right = nodeRightLeft.left
	nodeRightLeft.left = node
	nodeRightLeft.right = nodeRight

	node.subtreeHeight = NodeHeight(node)
	nodeRight.subtreeHeight = NodeHeight(nodeRight)
	nodeRightLeft.subtreeHeight = NodeHeight(nodeRightLeft)

	if root == node {
		root = nodeRightLeft
	}

	return nodeRightLeft // it has became the local root now
}

func RR_Rotation(node *Node) *Node {
	fmt.Println("performing RR rotation")
	nodeRight := node.right
	nodeRightLeft := nodeRight.left

	nodeRight.left = node
	node.right = nodeRightLeft

	node.subtreeHeight = NodeHeight(node)
	nodeRight.subtreeHeight = NodeHeight(nodeRight)

	if root == node {
		root = nodeRight
	}

	return nodeRight // it has became the local root now
}

func LR_Rotation(node *Node) *Node {
	fmt.Println("performing LR rotation")
	nodeLeft := node.left
	nodeLeftRight := nodeLeft.right

	nodeLeft.right = nodeLeftRight.left
	node.left = nodeLeftRight.right

	nodeLeftRight.left = nodeLeft
	nodeLeftRight.right = node

	node.subtreeHeight = NodeHeight(node)
	nodeLeft.subtreeHeight = NodeHeight(nodeLeft)
	nodeLeftRight.subtreeHeight = NodeHeight(nodeLeftRight)

	if root == node {
		root = nodeLeftRight
	}

	return nodeLeftRight // which has become the new local root
}

func LL_Rotation(node *Node) *Node {
	fmt.Println("performing LL rotation")
	nodeLeft := node.left
	nodeLeftRight := nodeLeft.right

	nodeLeft.right = node
	node.left = nodeLeftRight

	node.subtreeHeight = NodeHeight(node)
	nodeLeft.subtreeHeight = NodeHeight(nodeLeft)
	if root == node {
		root = nodeLeft
	}
	return nodeLeft // which has become the local root now

}

func BalanceFactor(node *Node) int {
	hleft := 0
	hright := 0
	//height of left and right subtree is alrady generated while inserting the nodes, just use the property form Node{}
	if node != nil && node.left != nil {
		hleft = node.left.subtreeHeight
	}
	if node != nil && node.right != nil {
		hright = node.right.subtreeHeight
	}

	return hleft - hright
}

func NodeHeight(node *Node) int {
	hleft := 0
	hright := 0
	if node != nil && node.left != nil {
		hleft = node.left.subtreeHeight
	}
	if node != nil && node.right != nil {
		hright = node.right.subtreeHeight
	}
	if hleft > hright {
		return hleft + 1
	} else {
		return hright + 1
	}
}

func PrintTree(w io.Writer, node *Node, spaceCount int, nodeType string) {
	if node != nil {
		for i := 0; i < spaceCount; i++ {
			fmt.Fprintf(w, " ")
		}
		fmt.Fprintf(w, "%v : %v (%v)\n", nodeType, node.data, node.subtreeHeight)
		PrintTree(w, node.left, spaceCount+2, "left")
		PrintTree(w, node.right, spaceCount+2, "rght")
	}
}

func InorderPredecessor(node *Node) *Node {
	if node == nil {
		return nil
	}
	if node != nil && node.right != nil {
		node = node.right
	}
	return node
}

func InorderSuccessor(node *Node) *Node {
	if node == nil {
		return nil
	}
	if node != nil && node.left != nil {
		node = node.left
	}
	return node
}

func RecursiveDelete(node *Node, key int) *Node {
	if node == nil {
		return nil
	}
	if node.left == nil && node.right == nil {
		if root == node {
			root = nil
		}
		return nil
	}
	if key < node.data {
		//call same function on the left node
		node.left = RecursiveDelete(node.left, key)
	} else if key > node.data {
		//call same function on the right node
		node.right = RecursiveDelete(node.right, key)
	} else if key == node.data {
		//node has been found
		//take inorder predecessor OR inorder successor based on the height of the subtree
		if NodeHeight(node.left) > NodeHeight(node.right) {
			//IN PRE
			inPre := InorderPredecessor(node.left)
			node.data = inPre.data
			node.left = RecursiveDelete(node.left, inPre.data)
		} else {
			//IN SUCC
			inSucc := InorderSuccessor(node.right)
			node.data = inSucc.data
			node.right = RecursiveDelete(node.right, inSucc.data)
		}
	}
	//now calculate the height of the node and perform rotations to balace the tree
	node.subtreeHeight = NodeHeight(node)
	//balance factor and rotation
	fmt.Println("BF node = ", BalanceFactor(node), " and BF node.left=", BalanceFactor(node.left), " and BF node.right=", BalanceFactor(node.right))
	if BalanceFactor(node) == 2 && BalanceFactor(node.left) == 1 {
		//LL rotation, L then L imbalance.
		return LL_Rotation(node)

	} else if BalanceFactor(node) == 2 && BalanceFactor(node.left) == -1 {
		//LR rrotation, L then R imbalance
		return LR_Rotation(node)

	} else if BalanceFactor(node) == -2 && BalanceFactor(node.right) == -1 {
		//RR rotation, R then R imbalance
		return RR_Rotation(node)
	} else if BalanceFactor(node) == -2 && BalanceFactor(node.right) == 1 {
		//RL rotation, R then L imbalance
		return RL_Rotation(node)
	}

	return node
}

func main() {
	// InsertIteratvie(50)
	// InsertIteratvie(40)
	// InsertIteratvie(60)
	// InsertIteratvie(55)
	// InsertIteratvie(58)
	// InsertIteratvie(70)
	// InsertIteratvie(30)
	// InsertIteratvie(38)
	// InsertIteratvie(35)
	// InsertIteratvie(42)
	// InsertIteratvie(33)
	root = InsertRecursive(root, 50)
	InsertRecursive(root, 40)
	InsertRecursive(root, 60)
	InsertRecursive(root, 55)
	InsertRecursive(root, 58)
	InsertRecursive(root, 70)
	InsertRecursive(root, 30)
	InsertRecursive(root, 58)
	InsertRecursive(root, 35)
	InsertRecursive(root, 42)
	InsertRecursive(root, 33)
	InsertRecursive(root, 43)
	InsertRecursive(root, 49)
	InsertRecursive(root, 46)
	InsertRecursive(root, 29)
	InsertRecursive(root, 25)
	InsertRecursive(root, 22)
	InsertRecursive(root, 15)
	InsertRecursive(root, 13)
	InsertRecursive(root, 11)
	InsertRecursive(root, 10)
	PrintTree(os.Stdout, root, 0, "root")
	RecursiveDelete(root, 43)
	PrintTree(os.Stdout, root, 0, "root")

}
