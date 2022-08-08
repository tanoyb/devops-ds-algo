package main

import "fmt"

//this is an example of backtracking code
var board [][]int = [][]int{
	{0, 0, 0, 0},
	{0, 0, 0, 0},
	{0, 0, 0, 0},
	{0, 0, 0, 0},
}

func solveNQueen(numOfRow int, board [][]int, currentRow int) int {
	//base case - return when the current row is equal to the num of rows of the board
	if currentRow == numOfRow {
		//print the board and return
		fmt.Println(board)
		return 1
	}

	//recurence relation/recursive case
	//try to place a queen in every row
	// i is the current row, so j is the columns of the row i

	ways := 0
	for j := 0; j < numOfRow; j++ {
		if canPlaceAQueen(board, numOfRow, currentRow, j) {
			//fill the position
			board[currentRow][j] = 1                           //if here we can place a queen, make recursive call to check if the remaining board has a success.
			ways += solveNQueen(numOfRow, board, currentRow+1) // if remaining board has a success, then we put the queen on position

			//if the remaining board is not a success, then we remove the queen position and try on the next position on the same row
			//backtrack
			board[currentRow][j] = 0

		}
	}
	//the above for loop is exited, that means we cant place the queen on that row in any position. we return false to the parent caller function.
	//then the parent caller function changses the queen's position on the next column of the row and checks again
	return ways
}

func canPlaceAQueen(board [][]int, numOfRows, x, y int) bool {
	//A queen moves left, right and diagonally, along those directions, there should be no other queens
	//column check
	for k := 0; k < x; k++ {
		if board[k][y] == 1 {
			return false
		}
	}

	//left diagonal
	i := x
	j := y
	for i >= 0 && j >= 0 {
		if board[i][j] == 1 {
			return false
		}
		i--
		j--
	}

	//right diagonal

	i = x
	j = y
	for i >= 0 && j < numOfRows {
		if board[i][j] == 1 {
			return false
		}
		i--
		j++
	}

	return true
}
func main() {
	fmt.Println(solveNQueen(4, board, 0))

}
