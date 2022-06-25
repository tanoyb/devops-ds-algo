package main

import "fmt"

func fib(n int) int {
	table := make([]int, n+1)
	table[1] = 1 //fib(1) is 1, so manually filling this up

	for i := 0; i <= n; i++ {
		if i+1 <= n {
			table[i+1] += table[i]
		}
		if i+2 <= n {
			table[i+2] += table[i]
		}
	}
	fmt.Println(table)

	return table[n]
}

func main() {
	fmt.Println(fib(50))
}
