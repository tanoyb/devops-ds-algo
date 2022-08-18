package main

import "fmt"

var count = 0
var memo map[int]int = map[int]int{}

//recrsive memoisation
func fib(n int) int {
	count++
	if val, ok := memo[n]; ok {
		return val
	}
	if n <= 2 {
		return 1
	}
	memo[n] = fib(n-1) + fib(n-2)
	return memo[n]
}

//tabulation method

func fib2(n int) int {
	//take an array as this is 1D DP
	arr := make([]int, n+1)
	arr[1], arr[2] = 1, 1
	for i := 3; i <= n; i++ {
		arr[i] = arr[i-1] + arr[i-2]
	}
	return arr[n]
}

func main() {
	fmt.Println(fib(50))
	fmt.Println("count = ", count)
	fmt.Println("tabulation = ", fib2(50))
}
