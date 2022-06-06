package main

import "fmt"

//the array have to be sorted,
//checks for element in the middle and splits the array
//three var lower, higher, mid=(lower+higher) / 2 floor value

func BinarySearch(arr []int, key int) int { //iterative version
	var low, mid, high int
	low = 0
	high = len(arr) - 1
	for low <= high {
		mid = (low + high) / 2 //takes floor value
		if key == arr[mid] {
			return mid
		} else if key < arr[mid] {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}

func BinarySearchRecursive(arr []int, low, high, key int) int { //recursive version
	var mid int
	if low <= high {
		mid = (low + high) / 2 //taking the floor value
		if key == arr[mid] {
			return mid
		} else if key < arr[mid] {
			return BinarySearchRecursive(arr, low, mid-1, key)
		} else {
			return BinarySearchRecursive(arr, mid+1, high, key)
		}
	}
	return -1
}

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	fmt.Println(BinarySearch(a, 11))
	fmt.Println(BinarySearchRecursive(a, 0, len(a), 11))

}
