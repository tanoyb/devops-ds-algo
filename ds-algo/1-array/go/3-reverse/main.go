package main

import "fmt"

//reversing an array using two methods.
//1st method is using an auxiliary array

//2nd method is using swapping techinique
// !!! Notice the use of pointers here in all the functions and
//passing a pointer of an array to other functions and
//passing the address of a pointer array to swap function

func ReverseArray(arr *[]int) {
	l := len(*arr) - 1
	//fmt.Println("l==", l)
	arr2 := make([]int, len(*arr))
	//fmt.Println("arr2", arr2)
	for i, j := l, 0; i >= 0; i, j = i-1, j+1 { //see the syntax for loop with two variables
		//fmt.Println("i=", i, ", j=", j)
		arr2[j] = (*arr)[i]
	}

	for i := 0; i < l; i++ {
		(*arr)[i] = arr2[i]
	}
	//fmt.Println("after change", *arr)
}

func swap(x *int, y *int) {
	*x, *y = *y, *x
}

func ReverseSwap(arr *[]int) {
	l := len(*arr) - 1
	for i, j := 0, l; i < j; i, j = i+1, j-1 {
		swap(&(*arr)[i], &(*arr)[j])
	}
}

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(a)
	ReverseArray(&a)
	fmt.Println(a)
	b := []int{11, 12, 13, 14, 15, 16, 17}
	ReverseSwap(&b)
	fmt.Println(b)
}

//for loop with two variable example
// func main() {
//     for i, j := 0, 1; i < 10; i, j = i+1, j+1 {
//         fmt.Println("Hello, playground")
//     }
// }
