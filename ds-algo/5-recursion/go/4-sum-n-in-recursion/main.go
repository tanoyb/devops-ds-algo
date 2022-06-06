package main

import "fmt"

func sum(n int) int {
	if n == 0 {
		return 0
	} else {
		return sum(n-1) + n
	}

}

func main() {
	s := sum(5)
	fmt.Println(s)

}
