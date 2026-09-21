package piscine

import "github.com/01-edu/z01"

func EightQueens() {
	var board [8]int
	solve(board, 0)
}

func solve(board [8]int, row int) {
	if row == 8 {
		printSolution(board)
		return
	}

	for col := 0; col < 8; col++ {
		if isSafe(board, row, col) {
			board[row] = col
			solve(board, row+1)
		}
	}
}

func isSafe(board [8]int, row, col int) bool {
	for i := 0; i < row; i++ {
		// Check if two queens are in the same column or diagonal
		if board[i] == col || board[i]-i == col-row || board[i]+i == col+row {
			return false
		}
	}
	return true
}

func printSolution(board [8]int) {
	for i := 0; i < 8; i++ {
		z01.PrintRune(rune(board[i] + '1'))
	}
	z01.PrintRune('\n')
}
