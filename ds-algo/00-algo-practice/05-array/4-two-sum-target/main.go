package main

import "fmt"

var arr = []int{2, 1, 5, 3}

var targetSum = 4

var previousHash = make(map[int]int) // value : index , it holds the elements which are checked in the loop already.

func main() {
	for i, v := range arr {
		diff := targetSum - v
		if _, ok := previousHash[diff]; ok {
			fmt.Println("target sum of ", targetSum, " is possible")
		} else {
			previousHash[v] = i
		}
	}
}
