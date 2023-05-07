package main

import "fmt"

//sudoku rules -> Each row must contain the numbers from 1 to 9, without repetitions
//Each column must contain the numbers from 1 to 9, without repetitions
//The digits can only occur once per subgrid(3*3 in size block) (nonet)
//The sum of every single row, column and nonet must equal 45

func printSudoku(board [][]int) {
	for i := 0; i < len(board); i++ {
		fmt.Println(board[i])
	}
}

func canPutNumber(board [][]int, i, j, num, n int) bool {
	//check for row and column if the number is present
	//n is 9 here
	for k := 0; k < n; k++ {
		if board[k][j] == num || board[i][k] == num {
			return false
		}
	}
	//check the subgrid/nonet if the bumber is alrady present
	sx := (i / 3) * 3
	sy := (j / 3) * 3
	for x := sx; x < sx+3; x++ {
		for y := sy; y < sy+3; y++ {
			if board[x][y] == num {
				return false
			}
		}
	}

	return true
}

func solveSudoku(board [][]int, i, j, n int) bool {
	//base case
	//n is the dimension, here the value of n is 9
	if i == n {

		//we reached to the end of the board of dimension n
		for d := 0; d < 9; d++ {
			for e := 0; e < 9; e++ {
				fmt.Printf("%d ", board[d][e])
			}
			fmt.Println()
		}
		return true
	}

	//recursive case
	if j == n {
		//reached the last column of the row i
		return solveSudoku(board, i+1, 0, n) // row i goes to next row i+1, column j starts again from 0 on the row i+1
	}
	//skip the prefilled cell
	if board[i][j] != 0 {
		//if the cell is not prefilled already
		return solveSudoku(board, i, j+1, n)
	}
	//cells to be filled with all possible number from 1 to 9
	for number := 1; number <= 9; number++ {
		if canPutNumber(board, i, j, number, n) {
			board[i][j] = number
			//now check if the rest of the subproblem can be solved
			IsSubProblemSolved := solveSudoku(board, i, j+1, n)
			if IsSubProblemSolved {
				return true
			}
		}
	}
	//if no option works
	board[i][j] = 0 //backtracking step
	return false
}

func main() {
	var sudokuBoard [][]int = [][]int{
		{5, 3, 0, 0, 7, 0, 0, 0, 0},
		{6, 0, 0, 1, 9, 5, 0, 0, 0},
		{0, 9, 8, 0, 0, 0, 0, 6, 0},
		{8, 0, 0, 0, 6, 0, 0, 0, 3},
		{4, 0, 0, 8, 0, 3, 0, 0, 1},
		{7, 0, 0, 0, 2, 0, 0, 0, 6},
		{0, 6, 0, 0, 0, 0, 2, 8, 0},
		{0, 0, 0, 4, 1, 9, 0, 0, 5},
		{0, 0, 0, 0, 8, 0, 0, 7, 9},
	}

	fmt.Println(solveSudoku(sudokuBoard, 0, 0, 9)) //n==9, is the dimension of the sudoku matrix board
}

//========================another style solution
/*

package main

import "fmt"

var sudokuBoard [][]int = [][]int{
	{5, 3, 0, 0, 7, 0, 0, 0, 0},
	{6, 0, 0, 1, 9, 5, 0, 0, 0},
	{0, 9, 8, 0, 0, 0, 0, 6, 0},
	{8, 0, 0, 0, 6, 0, 0, 0, 3},
	{4, 0, 0, 8, 0, 3, 0, 0, 1},
	{7, 0, 0, 0, 2, 0, 0, 0, 6},
	{0, 6, 0, 0, 0, 0, 2, 8, 0},
	{0, 0, 0, 4, 1, 9, 0, 0, 5},
	{0, 0, 0, 0, 8, 0, 0, 7, 9},
}

func solveSudoku(currentRow, currentColumn, matrixDimention int) bool {
	//base case will hit when we finish iterating all the rows
	if currentRow >= matrixDimention {
		fmt.Println()
		for i := 0; i < matrixDimention; i++ {
			fmt.Println(sudokuBoard[i])
		}
		return true
	}

	//recursive case
	//go to next row if end of a row is reaches
	if currentColumn > 8 {
		return solveSudoku(currentRow+1, 0, matrixDimention)
	}
	//skip if a position is already filled up
	if sudokuBoard[currentRow][currentColumn] != 0 {
		return solveSudoku(currentRow, currentColumn+1, matrixDimention)
	}

	//cells to be filled, we dont know what wil be the right number on this
	//so we try all possibilities
	for number := 1; number <= matrixDimention; number++ {
		if isSafeToPlaceNumber(number, currentRow, currentColumn) {
			sudokuBoard[currentRow][currentColumn] = number
			solveSubgrid := solveSudoku(currentRow, currentColumn+1, matrixDimention)
			if solveSubgrid {
				return true
			}

		}
	}
	//if no option work, we need to cal the previous call to check with next number
	sudokuBoard[currentRow][currentColumn] = 0 //reset

	return false
}

func isSafeToPlaceNumber(num, i, j int) bool {
	//check for row and column if the number is present
	//n is 9 here
	for k := 0; k < 9; k++ {
		if sudokuBoard[k][j] == num || sudokuBoard[i][k] == num {
			return false
		}
	}
	//check the subgrid/nonet if the bumber is alrady present
	sx := (i / 3) * 3
	sy := (j / 3) * 3
	for x := sx; x < sx+3; x++ {
		for y := sy; y < sy+3; y++ {
			if sudokuBoard[x][y] == num {
				return false
			}
		}
	}

	return true
}

func main() {
	fmt.Println(solveSudoku(0, 0, 9))
}


*/
