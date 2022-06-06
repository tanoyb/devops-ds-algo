package main

import "fmt"

func fib(n int) int { //using loop
	t0, t1 := 0, 1
	s := 0
	if n <= 1 {
		return 1
	} else {
		for i := 2; i <= n; i++ {
			s = t0 + t1
			t0 = t1
			t1 = s
		}
	}
	return s
}

func fibR(n int) int {
	if n <= 1 {
		return n
	}
	return fibR(n-2) + fibR(n-1)
}

func main() {
	fmt.Println(fib(5))
	fmt.Println(fibR(5))
}
