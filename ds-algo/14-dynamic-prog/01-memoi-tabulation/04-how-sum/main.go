package main

import "fmt"

//https://www.youtube.com/watch?v=oBt53YbR9Kk&t=5406s

//brute force
func howSum(targetSum int, nums []int) []int {
	if targetSum == 0 {
		return []int{}
	}
	if targetSum < 0 {
		return nil
	}
	for _, v := range nums {
		remainder := targetSum - v
		remainderResult := howSum(remainder, nums)
		if remainderResult != nil {
			a := []int{}
			a = append(a, remainderResult...)
			a = append(a, v)
			return a
		}

	}
	return nil
}

var memo = make(map[int][]int)

//DP memoization
func howSumDP(targetSum int, nums []int) []int {
	if _, ok := memo[targetSum]; ok {
		return memo[targetSum]
	}
	if targetSum == 0 {
		return []int{}
	}
	if targetSum < 0 {
		return nil
	}
	for _, num := range nums {
		remainder := targetSum - num
		remainderResult := howSum(remainder, nums)
		if remainderResult != nil {
			a := []int{}
			a = append(a, remainderResult...)
			a = append(a, num)
			memo[targetSum] = a
			return memo[targetSum]
		}

	}
	memo[targetSum] = nil
	return memo[targetSum]
}

func main() {
	fmt.Println(howSumDP(7, []int{2, 3}))
	fmt.Println(howSumDP(8, []int{2, 3, 5}))
	fmt.Println(howSumDP(300, []int{7, 14}))
}
