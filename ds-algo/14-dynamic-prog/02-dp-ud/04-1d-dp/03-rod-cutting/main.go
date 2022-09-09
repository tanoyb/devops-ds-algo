package main

import "fmt"

//		n=8		  // 1  2  3  4  5  6  7   8 //taking 1 based index rather than starting from 0
var price = []int{0, 1, 3, 4, 5, 7, 9, 10, 11}

//max piece of the rod 8
//add memoization

func solRodcutting(lengthOfPiece int) int {
	if lengthOfPiece == 0 {
		//base case
		return 0
	}
	ans := 0
	for i := 1; i <= lengthOfPiece; i++ {
		ans = getMax(ans, price[i]+solRodcutting(lengthOfPiece-i))
	}

	return ans
}
func getMax(n1, n2 int) int {
	if n1 > n2 {
		return n1
	} else {
		return n2
	}
}

func main() {
	fmt.Println(solRodcutting(8))
}
