package main

import (
	"fmt"
	"strings"
)

//brute force O(n^m)
func allConstruct(targetStr string, wordBank []string, wordIn string) [][]string {
	if targetStr == "" {
		return [][]string{
			{wordIn},
		}
	}
	var result [][]string

	for _, word := range wordBank {
		isPreFix := strings.HasPrefix(targetStr, word)
		if isPreFix == true {
			suffix := strings.TrimPrefix(targetStr, word)
			suffixways := allConstruct(suffix, wordBank, word)

			// if suffixways != nil {
			// 	fmt.Println("suffixways not nil", suffixways)
			// 	suffixways = append(suffixways, []string{word})
			// 	fmt.Println("found for word = ", word)
			// }
			for i, _ := range suffixways {
				//fmt.Println("suffixways before", suffixways[i], " and word is ", wordIn)
				suffixways[i] = append([]string{wordIn}, suffixways[i]...)
				//fmt.Println("suffixways[i] after", suffixways[i])
			}
			result = append(result, suffixways...)

		}
	}

	return result

}

var memo = make(map[string][][]string)

//emoization
func allConstructDP(targetStr string, wordBank []string, wordIn string) [][]string {
	if _, ok := memo[targetStr]; ok {
		return memo[targetStr]
	}

	if targetStr == "" {
		return [][]string{
			{wordIn},
		}
	}
	var result [][]string

	for _, word := range wordBank {
		isPreFix := strings.HasPrefix(targetStr, word)
		if isPreFix == true {
			suffix := strings.TrimPrefix(targetStr, word)
			suffixways := allConstructDP(suffix, wordBank, word)

			// if suffixways != nil {
			// 	fmt.Println("suffixways not nil", suffixways)
			// 	suffixways = append(suffixways, []string{word})
			// 	fmt.Println("found for word = ", word)
			// }
			for i, _ := range suffixways {
				//fmt.Println("suffixways before", suffixways[i], " and word is ", wordIn)
				suffixways[i] = append([]string{wordIn}, suffixways[i]...)
				//fmt.Println("suffixways[i] after", suffixways[i])
			}
			result = append(result, suffixways...)

		}
	}
	memo[targetStr] = result
	return result

}

func main() {
	fmt.Println(allConstruct("abcdef", []string{"ab", "abc", "cd", "def", "abcd", "ef", "c"}, ""))
	fmt.Println(allConstruct("purple", []string{"purp", "p", "ur", "le", "purpl"}, ""))
	fmt.Println(allConstruct("aabcdef", []string{"aab", "abc", "cd", "def", "abcd"}, ""))
	fmt.Println(allConstruct("tanoy", []string{"t", "tan", "cd", "an", "oy"}, ""))
	fmt.Println(allConstructDP("eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeef", []string{
		"e",
		"ee",
		"eee",
		"eeee",
		"eeeee",
	}, ""))
}
