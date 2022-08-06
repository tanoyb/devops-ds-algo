package main

import "fmt"

func Subset(str string, i int, curr string) {
	//base case
	if i == len(str) {
		fmt.Println(curr)
		return
	}

	//recurrence relation or recusive case

	//include the ith letter
	runeAtIndex := []rune(str)[i]      // in golang a character cant be taken using index. we convert it to []rune, and use index to ...
	charAtIndex := string(runeAtIndex) //pick the rune from [], and then convert that rune to a string character
	Subset(str, i+1, curr+charAtIndex)
	//exclude the ith letter
	Subset(str, i+1, curr)

}

func main() {
	Subset("abc", 0, "")

}

/*
func main() {
	s := "this is a stirng"
	s1 := []rune(s)
	fmt.Println(s1)
	for _, v := range s1 {
		fmt.Println(string(v)) // this line prints the actual characters from the string
	}
}
*/
