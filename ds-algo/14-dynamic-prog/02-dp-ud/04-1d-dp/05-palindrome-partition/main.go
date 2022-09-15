package main

import "fmt"

//- - - - -012345
var str = "aabbaabbaaccaaddaadddffgghhhssssaaakkkkllll"
var memo = map[string]bool{}
var palindromeCount = 0

//this is a brute force solution
//optimize this solution from On^2 to lower using memoization
func solvePalindromPartioning() {
	for i := 0; i < len(str); i++ {
		for j := i + 1; j < len(str); j++ {
			s := str[i : j+1]
			isPl := false
			if val, ok := memo[s]; ok {
				//result present in the hash table
				isPl = val
			} else {
				isPl = isPalindrome(s)
				memo[s] = isPl
			}
			if isPl {
				fmt.Println(s)
			}
		}
	}
	fmt.Println("length of the memo ", len(memo), " and plaindrome function was called ", palindromeCount, " time")
}

func isPalindrome(str string) bool {
	palindromeCount++
	reversedStr := ""
	for i := len(str) - 1; i >= 0; i-- {
		reversedStr += string(str[i])
	}
	for i := range str {
		if str[i] != reversedStr[i] {
			return false
		}
	}
	return true
}

func main() {
	solvePalindromPartioning()
}
