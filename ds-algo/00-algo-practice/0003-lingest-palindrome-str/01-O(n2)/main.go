package main

import (
	"fmt"
	"strings"
)

func longestPalindrome(str string) (string, []string) {
	var res string
	resLength := 0
	left := 0
	right := 0
	strArr := strings.Split(str, "")
	listArr := []string{}
	for i := 0; i < len(str); i++ {

		//odd length parts of the string where string length can be
		//1,3,5,7....
		left = i
		right = i
		for left >= 0 && right < len(strArr) && strArr[left] == strArr[right] {
			listArr = append(listArr, strings.Join(strArr[left:right+1], ""))
			if right-left+1 > resLength {
				res = strings.Join(strArr[left:right+1], "")
				resLength = right - left + 1
			}
			left--
			right++
		}

		//even length parts of the string where string length can be
		//2,4,6,8....
		left = i
		right = i + 1
		for left >= 0 && right < len(strArr) && strArr[left] == strArr[right] {
			listArr = append(listArr, strings.Join(strArr[left:right+1], ""))
			if right-left+1 > resLength {
				res = strings.Join(strArr[left:right+1], "")
				resLength = right - left + 1
			}
			left--
			right++
		}
	}
	return res, listArr
}

func main() {

	str := "aracecara"
	r, l := longestPalindrome(str)
	fmt.Println(r)
	if len(l) > 0 {
		for _, str := range l {
			fmt.Println(str)
		}
	}

}

//https://www.youtube.com/watch?v=XYQecbcd6_c&t=335s
