package main

import "fmt"

//this is an example of backtracking code
var board [][]int = [][]int{
	{0, 0, 0, 0},
	{0, 0, 0, 0},
	{0, 0, 0, 0},
	{0, 0, 0, 0},
}

func solveNQueen(numOfRow int, board [][]int, currentRow int) bool {
	//base case - return when the current row is equal to the num of rows of the board
	if currentRow == numOfRow {
		//print the board and return
		fmt.Println(board)
		return true
	}

	//recurence relation/recursive case
	//try to place a queen in every row
	// j is the columns of the row currentRow
	for j := 0; j < numOfRow; j++ {
		if canPlaceAQueen(board, numOfRow, currentRow, j) {
			//fill the position
			board[currentRow][j] = 1                              //if here we can place a queen, make recursive call to check if the remaining board has a success.
			success := solveNQueen(numOfRow, board, currentRow+1) // if remaining board has a success, then we put the queen on position
			if success {
				return true
			}
			//if the remaining board is not a success, then we remove the queen position and try on the next position on the same row
			//backtrack
			board[currentRow][j] = 0

		}
	}
	//the above for loop is exited, that means we cant place the queen on that row in any position. we return false to the parent caller function.
	//then the parent caller function changses the queen's position on the next column of the row and checks again
	return false
}

func canPlaceAQueen(board [][]int, numOfRows, currentRow, currentColumn int) bool {
	//A queen moves left, right and diagonally, along those directions, there should be no other queens
	//column check
	for k := 0; k < currentRow; k++ {
		if board[k][currentColumn] == 1 {
			return false
		}
	}

	//left diagonal
	i := currentRow
	j := currentColumn
	for i >= 0 && j >= 0 {
		if board[i][j] == 1 {
			return false
		}
		i--
		j--
	}

	//right diagonal

	i = currentRow
	j = currentColumn
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

//===========alternate code below
/*
package main

import "fmt"

var board [][]int = [][]int{
	{0, 0, 0, 0},
	{0, 0, 0, 0},
	{0, 0, 0, 0},
	{0, 0, 0, 0},
}

func solveNQueen(numRows int, currentRow int) bool {
	if currentRow == numRows {
		fmt.Println(board)
		return true
	}

	//iterate over the rows start from row 0 = current row
	for i := 0; i < numRows; i++ {
		if canPlaceAQueen(numRows, currentRow, i) {
			board[currentRow][i] = 1
			success := solveNQueen(numRows, currentRow+1)
			if success {
				return true
			}
			board[currentRow][i] = 0
		}
	}

	return false
}

func canPlaceAQueen(numOfRows, currentRow, currentColumn int) bool {
	//A queen moves left, right and diagonally, along those directions, there should be no other queens
	//column check
	for k := 0; k < currentRow; k++ {
		if board[k][currentColumn] == 1 {
			return false
		}
	}

	//left diagonal
	i := currentRow
	j := currentColumn
	for i >= 0 && j >= 0 {
		if board[i][j] == 1 {
			return false
		}
		i--
		j--
	}

	//right diagonal

	i = currentRow
	j = currentColumn
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
	r := solveNQueen(4, 0)
	if r {
		for i := 0; i < 4; i++ {
			fmt.Println(board[i])
		}
	}
}

*/
