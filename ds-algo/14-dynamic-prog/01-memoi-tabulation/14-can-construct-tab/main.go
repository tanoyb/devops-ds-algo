package main

import "fmt"

func canConstruct(target string, wordbank []string) bool {
	table := make([]bool, len(target)+1)
	table[0] = true //indicates that empty string can always be generated

	for i := 0; i <= len(target); i++ {
		if table[i] == true {
			for _, word := range wordbank {
				fmt.Println(word)
				incomplete
			}
		}

	}

	return false
}

func main() {
	fmt.Println(canConstruct("abcdef", []string{"ab", "abc", "cd", "def", "abcd"}))
}
