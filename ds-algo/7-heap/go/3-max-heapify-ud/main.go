package main

import "fmt"

//for insert into max heap heap, directoin of adjustment is "from leaf to root"
//for deleting root from max heap, directio of adjustment is "from root to leaf"
//Process of heapify :
//--1. First generate a complete binary tree
//--2. start from the last element in the array, that means
//the leaf element on the bottom right of the tree
//--3. the comparison will be with leaf of a element, that means
//we will compare the element with its children. when starting comparison
//from the last element(right most leaf), it has no children, so no comparison will be performed
//--4. the comparison wil start from last element (right most leaf of the tree), and the comparison will
//continue towards the left side leafs , and it will complete that level from right to left/
//once all the leafs are compared, then it will move to the upper level right most element
//and will follow the same procedure of compairson.
//--5. during this comparison the largest element will move to the root. once the largest element moves to the root(first element in the array)
//then, the previous root element will be compared with downward childrens to form the max heap(to satisfy the condition to be a max heap parent element >= child elemtns)
//trace the heapify function
// O(n)

func swap(A []int, i, j int) {
	A[i], A[j] = A[j], A[i]
}

func Delete(A []int, n int) int { //n is the size of tha array
	x := A[0] // Max element
	A[0] = A[n-1]

	i := 0
	j := 2*i + 1

	for j < n-1 {
		// Compare left and right children
		if A[j] < A[j+1] {
			j = j + 1
		}

		// Compare parent and largest child
		if A[i] < A[j] {
			swap(A, i, j)
			i = j
			j = 2*i + 1
		} else {
			break
		}
	}
	return x
}

func Heapify(A []int, n int) {
	//for a complete binary tree, #(number) of leaf elements: (n+1)/2, index of last leaf element's parent = (n/2)-1
	//as stated above, we will start from the bottom right leaf(last element in the array).
	for i := (n / 2) - 1; i >= 0; i-- {

		j := 2*i + 1 // Left child for current i

		for j < n-1 {
			// Compare left and right children of current i
			if A[j] < A[j+1] {
				j = j + 1
			}

			// Compare parent and largest child
			if A[i] < A[j] {
				swap(A, i, j)
				i = j
				j = 2*i + 1
			} else {
				break
			}
		}
	}
}

func main() {
	A := []int{5, 10, 30, 20, 35, 40, 15}
	fmt.Println(A)
	Heapify(A, len(A))
	fmt.Println(A)
}
