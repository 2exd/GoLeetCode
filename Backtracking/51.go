package main

import "strings"

func solveNQueens(n int) [][]string {
	board := make([][]string, n)
	for i := 0; i < n; i++ {
		board[i] = make([]string, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			board[i][j] = "."
		}
	}

	var res [][]string

	var backtracking func(row int)

	backtracking = func(row int) {
		if row == n {
			tmp := make([]string, n)
			for i := 0; i < n; i++ {
				tmp[i] = strings.Join(board[i], "")
			}
			res = append(res, tmp)
			return
		}

		for i := 0; i < n; i++ {

			if isValid(n, row, i, board) {
				board[row][i] = "Q"
				backtracking(row + 1)
				board[row][i] = "."
			}

		}

	}

	backtracking(0)

	return res
}

func isValid(n, row, col int, board [][]string) bool {
	// same row
	for i := 0; i < n; i++ {
		if board[i][col] == "Q" {
			return false
		}
	}

	// 45
	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if board[i][j] == "Q" {
			return false
		}
	}

	// 135
	for i, j := row-1, col+1; i >= 0 && j < n; i, j = i-1, j+1 {
		if board[i][j] == "Q" {
			return false
		}
	}
	return true
}
