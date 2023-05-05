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

/*
func GenerateBrakets(s string, n, currentIndex, countOpen, countClose int) {
	if currentIndex >= n*2 {
		if len(s) > 0 {
			Srune := []rune(s)
			firstChar := string(Srune[0])
			lastCar := string(Srune[len(s)-1])
			if firstChar == ")" || lastCar == "(" {
				return
			}

		}
		fmt.Println(s)
		return
	}
	//start with open
	if countOpen < n {
		GenerateBrakets(s+"(", n, currentIndex+1, countOpen+1, 0)
	}
	//with close
	if countClose < countOpen {
		GenerateBrakets(s+")", n, currentIndex+1, countOpen, countClose+1)
	}

}
*/
