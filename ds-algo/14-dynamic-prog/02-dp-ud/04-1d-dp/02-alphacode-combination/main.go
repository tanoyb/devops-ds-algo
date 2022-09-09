package main

import "fmt"

//var codeArr []int = []int{2, 5, 1, 1, 4}

var codeArr []int = []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
var memo map[int]int = map[int]int{}

func solAlphaCode(index int) int {
	if memo[index] != 0 {
		return memo[index]
	}
	if index == len(codeArr) {
		return 1
	}
	ans := 0
	if index < len(codeArr) && codeArr[index] >= 1 && codeArr[index] <= 9 { //considering single digit
		ans += solAlphaCode(index + 1)
	}
	if index < len(codeArr) && codeArr[index] == 1 { //considering code 10 to 19
		ans += solAlphaCode(index + 2)
	}
	if index < len(codeArr) && (codeArr[index] == 2 && codeArr[index+1] <= 6) { //considering code 20 to 26
		ans += solAlphaCode(index + 2)
	}
	memo[index] = ans
	return memo[index]
}

func main() {
	fmt.Println(solAlphaCode(0))
}
