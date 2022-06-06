package main

import "fmt"

//in bubble sort, in 1st pass, we can get the largest element,
//in 2nd pass, we can get the two largest sorted element.

func BubbleSort(arr []int, size int) {
	flag := 0 //we will use this to check if an array is already sorted
	for i := 0; i < size-1; i++ {
		flag = 0
		for j := 0; j < size-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j] //swap elements
				flag = 1
			}
		}
		if flag == 0 {
			//means no swaps, means array is already sorted
			break
		}
	}
}

func main() {
	A := []int{2, 5, 9, 8, 6, 5, 1, 56, 12, 23, 54, 87, 98, 32}
	BubbleSort(A, len(A))
	fmt.Println(A)
}
