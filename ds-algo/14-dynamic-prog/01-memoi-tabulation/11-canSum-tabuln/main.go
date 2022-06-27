package main

import "fmt"

func canSum(targetSum int, arr []int) bool {
	table := make([]bool, targetSum+1)
	table[0] = true //entry case or seed case
	for i := 0; i < len(table); i++ {
		fmt.Println("for main i=", i)
		if table[i] == true {
			for j := 0; j < len(arr); j++ {
				fmt.Println("for sub j=", j)
				fmt.Println("for cond i+arr[j] = ", i+arr[j], " and len table=", len(table))
				if i+arr[j] < len(table) {
					table[i+arr[j]] = true
				}
			}
		}
	}

	return table[targetSum]
}

func main() {

	fmt.Println("======", canSum(7, []int{5, 3, 4, 7}))
	fmt.Println("=======", canSum(7, []int{5, 4, 5}))
}
