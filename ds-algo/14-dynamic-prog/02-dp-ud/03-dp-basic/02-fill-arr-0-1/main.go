package main

import "fmt"

//calculate the number of ways to fill an array with 0 and 1 so that there are no consecutive 1s

//memoisation
func fillArr(currIdx, n int, number int) int {
	//n is the length of the array and isOne indicated if the value at index is 1 or not
	//base case
	if currIdx == n {
		return 1
	}

	answer := 0
	//if we place 0 first
	answer += fillArr(currIdx+1, n, 0)

	//if we place 1 first
	if number == 0 {
		answer += fillArr(currIdx+1, n, 1)
	}

	return answer
}

//tabulation

func main() {
	fmt.Println(fillArr(0, 3, 0))
}
