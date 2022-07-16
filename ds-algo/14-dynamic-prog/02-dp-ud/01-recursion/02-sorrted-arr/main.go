package main

import "fmt"

func isSorted(arr []int, start int) bool {
	if start == len(arr)-1 {
		return true
	}

	if arr[start] < arr[start+1] {
		return isSorted(arr, start+1)
	}

	return false
}
func main() {
	fmt.Println(isSorted([]int{1, 2, 3, 4, 5, 6}, 0))
	fmt.Println(isSorted([]int{1, 2, 3, 4, 6, 5}, 0))
}
