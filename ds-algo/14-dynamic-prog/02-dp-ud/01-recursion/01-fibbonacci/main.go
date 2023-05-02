package main

import "fmt"

func fib(n int) int {
	if n == 1 || n == 0 {
		return n
	}
	f1 := fib(n - 1)
	f2 := fib(n - 2)
	return f1 + f2
}

func main() {
	fmt.Println(fib(7))
}

//1 1 2 3 5 8 13 . . .
