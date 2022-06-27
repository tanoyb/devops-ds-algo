package main

import "fmt"

func canSum(targetSum int, arr []int) []int {
	table := make([][]int, targetSum+1)
	for i, _ := range table {
		table[i] = nil
	}
	table[0] = []int{} //entry case or seed case, when o is target, we return nil array
	for i := 0; i <= len(table)-1; i++ {
		//fmt.Println("for main i=", i)
		if table[i] != nil {
			for j := 0; j < len(arr); j++ {
				//fmt.Println("for sub j=", j)
				//fmt.Println("for cond i+arr[j] = ", i+arr[j], " and len table=", len(table))
				if i+arr[j] <= len(table)-1 {
					table[i+arr[j]] = mergeArr(table[i], arr[j])
				}
			}
			fmt.Println("table=", table)
		}
	}

	return table[targetSum]
}

func mergeArr(arr1 []int, n int) []int {
	//fmt.Println("merging==", arr1, arr2)
	arr := append(arr1, n)
	return arr
}

func main() {

	fmt.Println("======", canSum(7, []int{5, 3, 4}))
	fmt.Println("=======", canSum(7, []int{5, 4, 2}))
}
