package main

import (
	"fmt"
	"math"
)

//minimum steps to rach to 1
var memo map[int]int = map[int]int{}

func reachToOne(n int) int {
	if val, ok := memo[n]; ok {
		return val
	}
	if n == 1 {
		return 0
	}
	a := 999999999999
	b := 999999999999
	c := 999999999999
	a = 1 + reachToOne(n-1)
	if n%2 == 0 {
		b = 1 + reachToOne(n/2)
	}
	if n%3 == 0 {
		c = 1 + reachToOne(n/3)
	}
	memo[n] = int(math.Min(float64(a), math.Min(float64(b), float64(c))))
	return memo[n]
	//return int(math.Min(float64(a), math.Min(float64(b), float64(c))))
}

func main() {
	fmt.Println(reachToOne(5000))
}
