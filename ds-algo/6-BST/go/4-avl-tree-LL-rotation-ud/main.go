package main

import "fmt"

type Node struct {
	leftChild       *Node
	data            int
	heightOfSubtree int
	rightChild      *Node
}

var root *Node = nil

func InsertRecursive(node *Node, key int) *Node { //node is the root here
	if node == nil {
		temp := &Node{data: key, heightOfSubtree: 1}
		return temp
	}
	if key < node.data {
		node.leftChild = InsertRecursive(node.leftChild, key)
	} else if key > node.data {
		node.rightChild = InsertRecursive(node.rightChild, key)
	}

	node.heightOfSubtree = NodeHeight(node)
	//now calculate the balance factor of the inserted node
	if BalanceFactor(node) == 2 && BalanceFactor(node.leftChild) == 1 {
		// perform LL rotation because 'node' has +2 as balance factor
		//and 'node's left child has 1 as balance factor
		return LL_Rotation(node)
	} else if BalanceFactor(node) == 2 && BalanceFactor(node.leftChild) == -1 {
		return LR_Rotation(node)
	} else if BalanceFactor(node) == -2 && BalanceFactor(node.rightChild) == -1 {
		return RR_Rotation(node)
	} else if BalanceFactor(node) == -2 && BalanceFactor(node.rightChild) == 1 {
		return RL_Rotation(node)
	}

	return node
}

func LL_Rotation(p *Node) *Node {
	pl := p.leftChild
	plr := pl.rightChild

	pl.rightChild = p
	p.leftChild = plr
	p.heightOfSubtree = NodeHeight(p)
	pl.heightOfSubtree = NodeHeight(pl)

	if root == p {
		root = pl
	}
	return pl

}

func LR_Rotation(p *Node) *Node {
	pl := p.leftChild
	plr := pl.rightChild

	pl.rightChild = plr.leftChild
	p.leftChild = plr.rightChild

	plr.leftChild = pl
	plr.rightChild = p
	pl.heightOfSubtree = NodeHeight(pl)
	p.heightOfSubtree = NodeHeight(p)
	plr.heightOfSubtree = NodeHeight(plr)

	if root == p {
		root = plr
	}
	return plr
}

func RR_Rotation(node *Node) *Node {
	//own task
	return nil
}

func RL_Rotation(node *Node) *Node {
	//own task
	return nil
}

func BalanceFactor(node *Node) int {
	//height of left subtree - height of right subtree
	hleft := 0
	hright := 0
	if node != nil && node.leftChild != nil {
		hleft = node.leftChild.heightOfSubtree
	}
	if node != nil && node.rightChild != nil {
		hright = node.rightChild.heightOfSubtree
	}
	//if the balace factor is >= 2, then the node is imbalanced
	//if the return value is +2, then it is imbalanced on the left side
	//if the return value is -2, then it is imbalanced on the right side
	return hleft - hright
}

func NodeHeight(node *Node) int {
	hleft := 0
	hright := 0
	if node != nil && node.leftChild != nil {
		hleft = node.leftChild.heightOfSubtree
	}
	if node != nil && node.rightChild != nil {
		hright = node.rightChild.heightOfSubtree
	}
	if hleft > hright {
		return hleft + 1
	} else {
		return hright + 1
	}
}

func main() {
	//testing LL rotation
	// root = InsertRecursive(root, 10)
	// fmt.Println("====root==", root.data)
	// InsertRecursive(root, 5)
	// InsertRecursive(root, 2)

	// fmt.Println("root after rotation, root=", root.data)

	//----------
	//testing LR rotation
	root = InsertRecursive(root, 50)
	fmt.Println("====root==", root.data)
	InsertRecursive(root, 10)
	InsertRecursive(root, 20)
	fmt.Println("root after rotation, root=", root.data)

}
