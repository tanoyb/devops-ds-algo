package main

import "fmt"

var arr = []int{2, 7, 4, 5, 19}

//add memoization

func solveSubsetSum(currIdx, sum int) bool {
	//fmt.Println("curr idx = ", currIdx, " sum = ", sum)
	//base case
	if sum == 0 {
		return true
	}
	ans, ans1, ans2 := false, false, false
	//include the current index
	if currIdx < len(arr) && sum >= 0 {
		ans1 = solveSubsetSum(currIdx+1, sum-arr[currIdx])
		//fmt.Println("ans1=", ans1)

	}

	//exclude the current index
	if currIdx < len(arr) && sum >= 0 {
		ans2 = solveSubsetSum(currIdx+1, sum)
		//fmt.Println("ans2=", ans2)
	}

	if ans1 == true || ans2 == true {
		ans = true
	}
	return ans

}

func main() {
	fmt.Println(solveSubsetSum(0, 12))
}
