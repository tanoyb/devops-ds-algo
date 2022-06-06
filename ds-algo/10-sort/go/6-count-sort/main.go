package main

import "fmt"

var A = []int{6, 2, 4, 3, 9, 10, 6, 2, 2, 2, 0, 3, 8, 12, 6, 0, 0, 10, 17, 15, 1, 27, 23, 26}

func findMax() int {
	max := A[0]
	for i := 1; i < len(A); i++ {
		if A[i] > max {
			max = A[i]
		}
	}
	return max
}

func countSort() {
	max := findMax()
	countArr := make([]int, max+1)
	for i := 0; i < len(A); i++ {
		countArr[A[i]]++
	}
	fmt.Println("original array length=", len(A), " arr=", A)
	fmt.Println(countArr)
	//i is taken for aountarr looping and
	//j is taken to insert the values back to the array A
	//if we insert an element , we increment j by 1, otherwise we skip
	j := 0
	for i := 0; i < len(countArr); { //removed the i++ from here
		if countArr[i] > 0 {
			//fmt.Println("count arr index ", i, " has value ", countArr[i])
			A[j] = i      //start inserting elements back into A strting from index j=0
			j++           //after inserting an element, point j to the next index for next insert
			countArr[i]-- //reduce the countArr value by 1

		} else {
			i++ //once all the elements/duplicate elements are inserted into A, increment i for next index in count array.
		}
	}
	fmt.Println("sorted array length=", len(A), " arr=", A)
}

func main() {
	countSort()
}
