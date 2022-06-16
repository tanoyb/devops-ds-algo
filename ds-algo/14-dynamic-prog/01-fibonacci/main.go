package main

import "fmt"

//memoization (top down memoization)
var memo = make(map[int]int)

func fib(n int) int {
	if val, ok := memo[n]; ok {
		fmt.Println("found in memo")
		return val
	}
	if n <= 2 {
		return 1
	} else {
		memo[n] = fib(n-1) + fib(n-2)
		return memo[n]
	}
}

func main() {
	fmt.Println(fib(50))
}
