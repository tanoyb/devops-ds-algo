package main

import (
	"fmt"
	"strings"
)

//maintain the sequence of the words in targetStr

//brute force O(n^m)
func canConstruct(targetStr string, wordBank []string) bool {
	if targetStr == "" {
		return true
	}

	for _, word := range wordBank {
		isPreFix := strings.HasPrefix(targetStr, word)
		if isPreFix == true {
			suffix := strings.TrimPrefix(targetStr, word)
			if canConstruct(suffix, wordBank) == true {
				return true
			}
		}
	}

	return false

}

var memo = make(map[string]bool)

//memoization
func canConstructDP(targetStr string, wordBank []string) bool {
	if _, ok := memo[targetStr]; ok {
		return memo[targetStr]
	}
	if targetStr == "" {
		return true
	}

	for _, word := range wordBank {
		isPreFix := strings.HasPrefix(targetStr, word)
		if isPreFix == true {
			suffix := strings.TrimPrefix(targetStr, word)
			if canConstruct(suffix, wordBank) == true {
				memo[targetStr] = true
				return memo[targetStr]
			}
		}
	}
	memo[targetStr] = false
	return memo[targetStr]

}

func main() {
	fmt.Println(canConstruct("abcdef", []string{"ab", "abc", "cd", "def", "abcd"}))
	fmt.Println(canConstruct("aabcdef", []string{"aab", "abc", "cd", "def", "abcd"}))
	fmt.Println(canConstruct("tanoy", []string{"t", "tan", "cd", "an", "oy"}))
	fmt.Println(canConstruct("abcdefghijklmnop", []string{
		"a",
		"bc",
		"de",
		"efg",
		"dfg",
		"v",
		"t",
		"y",
		"i",
		"fgb",
		"er",
		"abcdefgh",
		"ijkl",
		"mnop",
	}))
}
