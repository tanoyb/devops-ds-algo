package main

import "fmt"

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

func main() {
	GenerateBrakets("", 4, 0, 0, 0)
}
