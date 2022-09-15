package main

import "fmt"

//add memoization to reduce time
//var amountArr = []int{1, 2, 3, 1} //correc result => 4
var amountArr = []int{2, 7, 9, 3, 1} //correct result => 12

func moneyRob(index int) int {
	fmt.Println("callinf for index ", index)
	if index > len(amountArr)-1 {
		fmt.Println("index out..returning ", index)
		return 0
	}
	amountEvenPosition := 0
	amountOddPosition := 0
	//start with index 0
	amountEvenPosition += amountArr[index] + moneyRob(index+2)
	//now consider the amount on next index, ad compare which one greater
	if index+1 < len(amountArr) {
		amountOddPosition += amountArr[index+1] + moneyRob(index+3)
	}
	if amountEvenPosition > amountOddPosition {
		return amountEvenPosition
	} else {
		return amountOddPosition
	}

}

func main() {
	fmt.Println("max amount can be robbed is ", moneyRob(0))
}
