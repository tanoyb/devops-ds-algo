package main

import "fmt"

func knapsack(profits, weights []int, capacity int, currentIndex int) int {
	//base case
	if capacity <= 0 || currentIndex >= len(profits) {
		return 0
	}

	profit := 0
	// recursive call after choosing the element at the currentIndex
	// if the weight of the element at currentIndex exceeds the capacity, we shouldn't process this
	if weights[currentIndex] <= capacity {
		profit = profits[currentIndex] + knapsack(profits, weights, capacity-weights[currentIndex], currentIndex+1)
	}
	// recursive call after excluding the element at the currentIndex
	profit2 := knapsack(profits, weights, capacity, currentIndex+1)

	if profit > profit2 {
		return profit
	} else {
		return profit2
	}
}

func main() {
	profits := []int{1, 6, 10, 16}
	weights := []int{1, 2, 3, 5}
	fmt.Println(knapsack(profits, weights, 7, 0))
}

/*
Given the weights and profits of ‘N’ items, we are asked to put these items in a knapsack that has a capacity ‘C’. The goal is to get the maximum profit from the items in the knapsack. Each item can only be selected once, as we don’t have multiple quantities of any item.

Let’s take Merry’s example, who wants to carry some fruits in the knapsack to get maximum profit. Here are the weights and profits of the fruits:

Items: { Apple, Orange, Banana, Melon }
Weights: { 2, 3, 1, 4 }
Profits: { 4, 5, 3, 7 }
Knapsack capacity: 5
*/
