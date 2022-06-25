package main

import (
	"fmt"
	"strings"
)

//brute force O(n^m)
func countConstruct(targetStr string, wordBank []string) int {
	if targetStr == "" {
		return 1
	}

	total := 0

	for _, word := range wordBank {
		isPreFix := strings.HasPrefix(targetStr, word)
		if isPreFix == true {
			suffix := strings.TrimPrefix(targetStr, word)
			numWaysForTheRest := countConstruct(suffix, wordBank)
			total += numWaysForTheRest
		}
	}

	return total

}

var memo = make(map[string]int)

//DP memo
func countConstructDP(targetStr string, wordBank []string) int {
	if val, ok := memo[targetStr]; ok {
		return val
	}
	if targetStr == "" {
		return 1
	}

	totalCount := 0

	for _, word := range wordBank {
		isPreFix := strings.HasPrefix(targetStr, word)
		if isPreFix == true {
			suffix := strings.TrimPrefix(targetStr, word)
			numWaysForTheRest := countConstructDP(suffix, wordBank)
			totalCount += numWaysForTheRest
		}
	}
	memo[targetStr] = totalCount

	return memo[targetStr]

}

func main() {
	fmt.Println(countConstructDP("abcdef", []string{"ab", "abc", "cd", "def", "abcd"}))
	fmt.Println(countConstructDP("aabcdef", []string{"aab", "abc", "cd", "def", "abcd"}))
	fmt.Println(countConstructDP("tanoy", []string{"t", "tan", "cd", "an", "oy"}))
	fmt.Println(countConstructDP("eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeef", []string{
		"e",
		"ee",
		"eee",
		"eeee",
		"eeeee",
	}))
}
