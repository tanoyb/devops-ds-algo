package main

import (
	"fmt"
	"strings"
)

//another approach for this problen is sorting.
//if we sort these two strings, they should be exactly equal
//then we can decide that whether they are anagram or not

var s = "anagram"
var t = "nagaram"

//check if t is an anagram of s
var mapS = make(map[string]int)
var mapT = make(map[string]int)

func GenerateHashMap(arr []string, mapName *map[string]int) {
	for _, s := range arr {
		(*mapName)[s] += 1
	}
	fmt.Println((*mapName))
}

func main() {
	arrS := strings.Split(s, "")
	arrT := strings.Split(t, "")
	flag := true

	if len(s) != len(t) {
		flag = false
	} else {
		GenerateHashMap(arrS, &mapS)
		GenerateHashMap(arrT, &mapT)

		for str, _ := range mapT {
			if mapT[str] != mapS[str] {
				flag = false
				break
			}
		}
	}

	if flag {
		fmt.Println("yes an anagram")
	} else {
		fmt.Println("not an anagram")
	}
}
