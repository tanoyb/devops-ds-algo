package main

import "fmt"

// take the first element and find the minimum element in the list by comparing with the first element
//swap the first element with the minimum element.
//on the 1st pass, we will get the minimum element in the list

func SelectionSort(A []int, n int) {
	i := 0
	j := 0
	k := 0
	for i = 0; i < n-1; i++ { // i indicates the number of pass
		//now this 2nd for loop is for comparing element in rach pass
		for j, k = i, i; j < n; j++ {
			if A[j] < A[k] {
				k = j
			}
		}
		A[i], A[k] = A[k], A[i]
	}
}

func main() {
	A := []int{2, 5, 9, 8, 6, 5, 1, 56, 12, 23, 54, 87, 98, 32}
	//BubbleSort(A, len(A))
	//fmt.Println(A)
	SelectionSort(A, len(A))
	fmt.Println(A)
}
