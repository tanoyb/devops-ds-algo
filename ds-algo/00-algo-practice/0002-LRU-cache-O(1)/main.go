package main

import (
	"fmt"
	"time"
)

var cacheSize = 4

type Node struct {
	left  *Node
	data  int
	right *Node
}

type Cache struct {
	length int
	front  *Node
	rear   *Node
}

var CacheMap = make(map[int]*Node)

//rules for the cache is mentioned on the schrrenshots and video
//considering cache length is 3
func insertIntoCache(c *Cache, n int) {
	node := searchInCache(c, n)
	if node != nil {
		fmt.Println("searched element present in cache..updating positions")
		fmt.Println("searched node ", node)
		if node == c.front {
			//matching with the first node. in this case, make the first node
			//as last node and put every node one step forward
			fmt.Println("matching with front node..moving to last....moving each node one step forward")
			tempFront := c.front
			//change front
			c.front = c.front.right
			c.front.left = nil
			//change rear
			c.rear.right = tempFront
			tempFront.left = c.rear
			c.rear = tempFront
			c.rear.right = nil

		} else if node == c.rear {
			//mathing with last node. no action is required
			fmt.Println("matching with rear node..no actions required")

		} else {
			//mathing with some middle elements, swap positions
			fmt.Println("matching with middle node..adding at last...moving each node one step forward")
			prevNode := node.left
			nextNode := node.right
			tempRear := c.rear
			prevNode.right = nextNode
			nextNode.left = prevNode
			c.rear = node
			c.rear.left = tempRear
			tempRear.right = c.rear
		}

	} else {
		fmt.Println("search element not in cache..adding it")
		if c.length == 0 {
			fmt.Println("empty cache..creating first node")
			newNode := &Node{left: nil, data: n, right: nil}
			c.front = newNode
			c.rear = newNode
			c.length++
			CacheMap[n] = newNode
		} else if c.length >= 3 {
			fmt.Println("cache is full..replacing removing old from front")
			newNode := &Node{left: nil, data: n, right: nil}
			//move front
			CacheMap[c.front.data] = nil
			c.front = c.front.right
			c.front.left = nil
			//now move rear
			newNode.left = c.rear
			c.rear.right = newNode
			c.rear = newNode
			CacheMap[n] = newNode
		} else {
			fmt.Println("space available...inserting into cache")
			newNode := &Node{left: nil, data: n, right: nil}
			c.rear.right = newNode
			newNode.left = c.rear
			c.rear = newNode
			c.length++
			CacheMap[n] = newNode
		}
	}

}

func searchInCache(c *Cache, n int) *Node {
	fmt.Println("searching for ", n)
	if node, ok := CacheMap[n]; ok {
		fmt.Println("search successful, found node ", node)
		return node
	}

	return nil
}

func printCache(c *Cache) {
	toPrint := c.front
	length := c.length
	for length != 0 {
		fmt.Println(toPrint.data)
		toPrint = toPrint.right
		length--
	}
}

func main() {
	timestart := time.Now()
	//pageNum := []int{1, 2, 3, 4, 5, 3, 6, 7, 3, 4, 5, 2, 1}

	cacheLL := &Cache{length: 0, front: nil, rear: nil}
	insertIntoCache(cacheLL, 1)
	insertIntoCache(cacheLL, 2)
	insertIntoCache(cacheLL, 3)
	//insertIntoCache(cacheLL, 4)
	//insertIntoCache(cacheLL, 5)
	fmt.Println("printing cache length = ", cacheLL.length)
	printCache(cacheLL)

	insertIntoCache(cacheLL, 2)
	printCache(cacheLL)
	insertIntoCache(cacheLL, 1)
	printCache(cacheLL)
	insertIntoCache(cacheLL, 3)
	printCache(cacheLL)
	fmt.Println("total time taken = ", time.Since(timestart))

}

//https://www.youtube.com/watch?v=akFRa58Svug
