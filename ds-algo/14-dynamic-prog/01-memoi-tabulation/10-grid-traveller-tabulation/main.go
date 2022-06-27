package main

import "fmt"

//same problem using tabulation,
//please look at the grid traveller screenshots on memoization section
//https://www.youtube.com/watch?v=oBt53YbR9Kk&t=12279s

func gridTraveller(m, n int) int {
	table := make([][]int, m+1)
	for i := range table {
		table[i] = make([]int, n+1)
	}
	//fmt.Println(table)
	table[1][1] = 1
	for i := 0; i <= m; i++ {
		for j := 0; j <= n; j++ {
			current := table[i][j]
			//add to the right
			if j+1 <= n {
				table[i][j+1] += current
			}

			//add to the down
			if i+1 <= m {
				table[i+1][j] += current
			}
		}
	}
	return table[m][n]

}

func main() {
	fmt.Println(gridTraveller(3, 2))
	fmt.Println(gridTraveller(18, 18))
}
