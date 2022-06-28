package main

import "fmt"

//https://www.youtube.com/watch?v=oBt53YbR9Kk&t=14890s

func bestSum(targetSum int, numbers []int) []int {
	table := make([][]int, targetSum+1)
	for i, _ := range table {
		table[i] = nil
	}
	table[0] = []int{} //entry case or seed case, when o is target, we return nil array
	for i := 0; i <= len(table)-1; i++ {
		//fmt.Println("for main i=", i)
		if table[i] != nil {
			for j := 0; j < len(numbers); j++ {
				//fmt.Println("for sub j=", j)
				//fmt.Println("for cond i+arr[j] = ", i+arr[j], " and len table=", len(table))
				if i+numbers[j] <= len(table)-1 {
					//fmt.Println("==current arr = ", table[i+numbers[j]])
					//fmt.Println("==arr to merge=", mergeArr(table[i], numbers[j]))
					combination := mergeArr(table[i], numbers[j])
					if table[i+numbers[j]] == nil || len(table[i+numbers[j]]) > len(combination) {

						table[i+numbers[j]] = combination
					}

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

	fmt.Println("======", bestSum(8, []int{2, 3, 5}))
	//fmt.Println("=======", howSum(7, []int{5, 4, 2}))
}
