package main

import "fmt"

//the solution is discussed on the screenshot

func tilingWays(n int) int {
	if n <= 3 {
		return 1
	}
	return tilingWays(n-1) + tilingWays(n-4)
}

func main() {
	fmt.Println(tilingWays(10))
}
