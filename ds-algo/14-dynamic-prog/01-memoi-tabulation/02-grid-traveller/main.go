package main

import "fmt"

//https://www.youtube.com/watch?v=oBt53YbR9Kk&list=PLhwld_MawMmHTqDFmrZZeZ3W1yf5cJcd2&index=1&t=2489s

var memo = make(map[string]int)

func gridTraveller(m, n int) int {
	key := fmt.Sprintf("%d,%d", m, n)
	if val, ok := memo[key]; ok {
		fmt.Println("found in memo")
		return val
	}
	if m == 1 && n == 1 {
		return 1
	}
	if m == 0 || n == 0 {
		return 0
	}
	//now go down and then go right and add the results
	//when going down reduce row by 1 and when going right reduce column by 1
	memo[key] = gridTraveller(m-1, n) + gridTraveller(m, n-1)
	return memo[key]
}

func main() {
	fmt.Println(gridTraveller(18, 18))
}
