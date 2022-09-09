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

//tabulation

func reachToOneTabuln(N int) int {
	dp := make([]int, N+1)

	dp[0] = 0
	dp[1] = 0
	dp[2] = 1
	dp[3] = 1

	for i := 4; i <= N; i++ {
		a := 9999999999999999
		b := 99999999999999999
		c := 99999999999999999

		if i%3 == 0 {
			a = 1 + dp[i/3]
		}

		if i%2 == 0 {
			b = 1 + dp[i/2]
		}

		c = 1 + dp[i-1]
		dp[i] = int(math.Min(float64(a), math.Min(float64(b), float64(c))))
	}

	return dp[N]
}

func main() {
	fmt.Println(reachToOne(5000))
}
