package main

import "fmt"

func GenerateBrakets(output string, n, open, close, currentIndex int) {
	//n is the length of the string, if n = 2, then length should be 2*2=4, because two '(' is there then there has to be two ')' to balacne it.
	//thats the reason the actual length of the output string will be 2 * n, if n '(', then we need n ')', total 2n.
	//base case
	if currentIndex >= 2*n {
		fmt.Println(output)
		return
	}
	//open
	if open < n {
		GenerateBrakets(output+"(", n, open+1, close, currentIndex+1)

	}
	//close breakets
	if close < open {
		GenerateBrakets(output+")", n, open, close+1, currentIndex+1)
	}
}

func main() {
	GenerateBrakets("", 3, 0, 0, 0)
}
