package main

import "fmt"

//simple binary search algo

var nums []int = []int{6, 7, 8, 9, 10, 11, 12, 13, 123, 234, 345, 456, 567, 678}

func main() {
	target := 8
	low := 0
	high := len(nums) - 1
	mid := 0
	for low <= high {
		mid = (low + high) / 2
		if target > nums[mid] {
			//search on right
			low = mid + 1
		} else if target < nums[mid] {
			//search on left
			high = mid - 1
		} else {
			fmt.Println("element found at index ", mid)
			break
		}
	}

}
