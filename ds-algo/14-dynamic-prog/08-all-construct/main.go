package main

import (
	"fmt"
	"strings"
)

//brute force O(n^m)
func allConstruct(targetStr string, wordBank []string) [][]string {
	if targetStr == "" {
		return [][]string{}
	}
	var result [][]string

	for _, word := range wordBank {
		isPreFix := strings.HasPrefix(targetStr, word)
		if isPreFix == true {
			suffix := strings.TrimPrefix(targetStr, word)
			suffixways := allConstruct(suffix, wordBank)

			fmt.Println("returned from func", suffixways)
			if len(suffixways) == 0 {
				suffixways = append(suffixways, []string{word})
			}
			for i, _ := range suffixways {
				fmt.Println("suffixways before", suffixways[i])
				suffixways[i] = append([]string{word}, suffixways[i]...)
				fmt.Println("suffixways[i] after", suffixways[i])
			}
			result = append(result, suffixways...)

		}
	}

	return result

}
still has problems

func main() {
	fmt.Println(allConstruct("abcdef", []string{"ab", "abc", "cd", "def", "abcd"}))
	//fmt.Println(allConstruct("aabcdef", []string{"aab", "abc", "cd", "def", "abcd"}))
	//fmt.Println(allConstruct("tanoy", []string{"t", "tan", "cd", "an", "oy"}))
	// fmt.Println(allConstruct("eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeef", []string{
	// 	"e",
	// 	"ee",
	// 	"eee",
	// 	"eeee",
	// 	"eeeee",
	// }))
}
