package main

import "fmt"

var allArray [][]int

func bestSum(targetSum int, nums []int) []int { //brute force
	if targetSum == 0 {
		return []int{}
	}
	if targetSum < 0 {
		return nil
	}
	for _, number := range nums {
		remainder := targetSum - number
		returnArr := bestSum(remainder, nums)
		if returnArr != nil {
			var combination []int
			combination = append(combination, returnArr...)
			combination = append(combination, number)
			// if len(shortestCombination) == 0 || len(combination) < len(shortestCombination) {
			// 	shortestCombination = combination
			// }
			return combination
		}
	}

	return nil
}

//write the memoization approach

func main() {
	fmt.Println(bestSum(8, []int{5, 3, 2}))
}
