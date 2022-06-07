package main

import "fmt"

//sorted array and conditions are in the diagram.
//brute force approach is O(n^2)
//O(n) approach:
//check the max element index in the sorted arrya which is equals to grater than the target sum.
//consider the array upto that 0 to index which was found on above condition.
//set left position to 0 index and right position to that calculated index.
//check the sum, if sum > target, reduce the right position
//if sum < target, increase the left position, at one position, the correct indexes will be found

var arr = []int{1, 3, 4, 5, 7, 10, 11, 13, 17, 21, 25}
var target = 13

func BinarySearchRecursive(low, high int) int { //perform binary search to get the element
	//a := math.Floor(7 / 2)//example
	//x := math.Ceil(1.36)
	mid := (low + high) / 2

	fmt.Println("low=", low, " high=", high, " mid=", mid)
	if arr[mid] == target {
		return mid
	}
	if low == high || low >= high {
		return -1
	}
	if arr[mid] > target {
		return BinarySearchRecursive(low, mid-1)
	} else {
		return BinarySearchRecursive(mid+1, high)
	}
}

func getUpperLimitIndex(low, high int) int {
	//taken from above binary search example

	if arr[high] <= target {
		return high
	} else if arr[low] > target {
		return -1
	}
	var index = -1
	for low < high {
		mid := (low + high) / 2
		fmt.Println("low=", low, " mid=", mid, " high=", high)
		if arr[mid] == target {
			index = mid
			break
		}
		if arr[mid] < target {
			low = mid + 1
		} else if arr[mid] > target {
			high = mid
			index = high
			break
		}

	}
	return index
}

func main() {
	//first check that index which value is >= target
	maxIndex := getUpperLimitIndex(0, len(arr)-1)
	fmt.Println(maxIndex)
	//now compare the sum to be 9
	takenIndexes := [50]int{}
	low := 0
	high := maxIndex
	flag := false
	for low <= high {
		fmt.Println("checking..........low=", low, " high=", high)
		sum := arr[low] + arr[high]
		if high > maxIndex {
			break
		}
		if sum == target {
			flag = true
			if takenIndexes[low] == 0 && takenIndexes[high] == 0 {
				fmt.Println("the desired indexes are ", low, high)

			}
			takenIndexes[low] = 1
			takenIndexes[high] = 1
			low = low + 1
			high = high - 1 //because low and high indexes are already taken, so cant use them twice

		} else if sum > target {
			//reduce the high by 1, because if we hike low by 1, the sum will be greater
			high = high - 1
		} else if sum < target {
			//hike low by 1
			low = low + 1
		}
	}
	if flag {
		fmt.Println("results shown")
	} else {
		fmt.Println("no match")
	}
}

//https://www.youtube.com/watch?v=cQ1Oz4ckceM&list=WL&index=14
