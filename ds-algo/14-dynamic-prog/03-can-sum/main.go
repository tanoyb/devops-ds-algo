package main

import "fmt"

//https://www.youtube.com/watch?v=oBt53YbR9Kk&t=3726s

/*func canSum(targetSum int, arr []int) bool {//BRUTE FORCE
	if targetSum == 0 {
		return true
	}
	if targetSum < 0 {
		return false
	}
	for _, v := range arr {
		fmt.Println("num=", v)
		remainder := targetSum - v
		if canSum(remainder, arr) == true {
			return true
		}
	}
	return false

}*/

var memo = make(map[int]bool)

func canSum(targetSum int, arr []int) bool { //memoisation
	if _, ok := memo[targetSum]; ok {
		return memo[targetSum]
	}
	if targetSum == 0 {
		return true
	}
	if targetSum < 0 {
		return false
	}
	for _, v := range arr {
		fmt.Println("num=", v)
		remainder := targetSum - v

		if canSum(remainder, arr) == true {
			memo[targetSum] = true
			return true
		}
	}
	memo[targetSum] = false
	return false

}

func main() {
	fmt.Println(canSum(7, []int{5, 3, 4, 7}))
	fmt.Println(canSum(300, []int{7, 14}))
}
