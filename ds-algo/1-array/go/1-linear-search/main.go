package main

import "fmt"

//improved linear search by doing transposition/move to front by 1 index
func swapWithPreviousIndex(arr *[]int, i int) {
	if i > 0 {
		(*arr)[i-1], (*arr)[i] = (*arr)[i], (*arr)[i-1]
	}
}

func LinearSearch(arr *[]int, key int) int {
	//fmt.Println("len==", len(*arr))
	//fmt.Println("element==", (*arr)[0])
	for i := 0; i < len(*arr); i++ {
		if key == (*arr)[i] {
			//swap to improve
			swapWithPreviousIndex(arr, i)
			return i
		}
	}
	return -1
}

func main() {
	arr := []int{2, 3, 4, 5, 6, 7, 8, 9, 11, 12, 13, 234, 34, 54, 65}
	fmt.Println(arr)
	fmt.Println(&arr)
	index := LinearSearch(&arr, 65)
	LinearSearch(&arr, 65)
	LinearSearch(&arr, 65)
	LinearSearch(&arr, 65) //improvement by shifting the elemnt to fornt of the array by 1 index
	fmt.Println("element index=", index)
	fmt.Println(arr)
}
