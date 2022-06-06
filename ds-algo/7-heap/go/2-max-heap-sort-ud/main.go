package main

import "fmt"

//heap is a complete binary tree
//stored in an array
//only root can be deleted from a heap
//max heap -- a parent element >= all decendenats elements
//min heap -- a parent element <= all decendenats elements
//used in hea sort technique

//heap sort : 1. create heap of n elements - O(nlogn)
//2. delete n elements one by one - O(nlogn)

//left child index= [parent index]*2 + 1
//right child index = [parent index]*2 + 2

//index of parent = child index / 2 => floor value

func InsertIntoMaxHeap(heapArr []int, i int) { //i is the index position
	temp := heapArr[i]

	for i > 1 && temp > heapArr[i/2] { //taking floor value by default
		heapArr[i] = heapArr[i/2]
		i = i / 2
	}
	heapArr[i] = temp //because value of i will be i/2 now
}

func DeleteFromHeapTop(arr []int, n int) int { //n is the array size
	//[]int{-1, 10, 20, 30, 25, 5, 40, 35}
	temp := 0
	val := arr[1]
	arr[1] = arr[n]
	arr[n] = val
	i := 1
	j := i * 2

	for j < n-1 { //2 arr size is decremented by 1, as we deleted last element from the array
		if arr[j+1] > arr[j] {
			j = j + 1
		}
		if arr[i] < arr[j] {
			temp = arr[i]
			arr[i] = arr[j]
			arr[j] = temp
			i = j
			j = 2 * j
		} else {
			break
		}
	}
	return val
}

func main() {
	HeapArr := []int{-1, 10, 20, 30, 25, 5, 40, 35}
	//first element as -1, because we will consider here that
	//array index starting from 1, for ease of formula calculation
	for i := 2; i < len(HeapArr); i++ {
		InsertIntoMaxHeap(HeapArr, i)
	}
	fmt.Println(HeapArr)
	//fmt.Printf("deleting value %d\n", DeleteFromHeapTop(HeapArr, 7))
	//fmt.Println(HeapArr)
	for i := 7; i > 1; i-- {
		fmt.Println(DeleteFromHeapTop(HeapArr, i))
	}

}
