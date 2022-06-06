package main

import "fmt"

func fun(n int) int {
	if n > 100 {
		return n - 10
	} else {
		return fun(fun(n + 11))
	}
}

func main() {
	r := fun(95)
	fmt.Println(r)
}
