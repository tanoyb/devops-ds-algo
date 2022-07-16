package main

import "fmt"

func increasing(n int) {
	if n > 0 {
		increasing(n - 1)
		fmt.Println(n)

	}

}

func decreasing(n int) {
	if n > 0 {
		fmt.Println(n)
		decreasing(n - 1)
	}
}

func main() {
	a := 5
	increasing(a)
	decreasing(a)
}
