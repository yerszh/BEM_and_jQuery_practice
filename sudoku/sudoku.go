package main

import (
	"fmt"
	"os"
)

func main() {
	board := os.Args[1:]

	if !isValidBoard(board) {
		fmt.Println("Error")
		return
	}

	solved := solveSudoku(board)

	if solved {
		printBoard(board)
	} else {
		fmt.Println("Error")
	}
}

func solveSudoku(board []string) bool {
	row, col := 0, 0
	isEmpty := true

	// Find the first empty cell in the board
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				row, col = i, j
				isEmpty = false
				break
			}
		}
		if !isEmpty {
			break
		}
	}

	if isEmpty {
		return true // All cells are filled, puzzle solved
	}

	// Try each digit from 1 to 9
	for num := 1; num <= 9; num++ {
		if isSafe(board, row, col, num) {
			// Place the number in the cell
			board[row] = replaceCharAtIndex(board[row], col, byte(num)+'0')

			if solveSudoku(board) {
				return true
			}

			// If the number doesn't lead to a solution, undo the placement
			board[row] = replaceCharAtIndex(board[row], col, '.')
		}
	}

	return false
}

func isSafe(board []string, row, col, num int) bool {
	// Check if the number is already present in the current row
	for i := 0; i < 9; i++ {
		if board[row][i] != '.' && board[row][i]-'0' == byte(num) {
			return false
		}
	}

	// Check if the number is already present in the current column
	for i := 0; i < 9; i++ {
		if board[i][col] != '.' && board[i][col]-'0' == byte(num) {
			return false
		}
	}

	// Check if the number is already present in the 3x3 box
	startRow, startCol := (row/3)*3, (col/3)*3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[startRow+i][startCol+j] != '.' && board[startRow+i][startCol+j]-'0' == byte(num) {
				return false
			}
		}
	}

	return true
}

func replaceCharAtIndex(s string, index int, char byte) string {
	runes := []byte(s)
	runes[index] = char
	return string(runes)
}

func printBoard(board []string) {
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			fmt.Print(string(board[row][col]))
			if col < 8 {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func isValidBoard(board []string) bool {
	if len(board) != 9 {
		return false
	}

	// Check for duplicate values in rows
	for i := 0; i < 9; i++ {
		if len(board[i]) != 9 {
			return false
		}

		seen := make(map[byte]bool)
		for j := 0; j < 9; j++ {
			cell := board[i][j]
			if cell != '.' {
				if cell < '1' || cell > '9' || seen[cell] {
					return false
				}
				seen[cell] = true
			}
		}
	}

	// Check for duplicate values in columns
	for j := 0; j < 9; j++ {
		seen := make(map[byte]bool)
		for i := 0; i < 9; i++ {
			cell := board[i][j]
			if cell != '.' {
				if cell < '1' || cell > '9' || seen[cell] {
					return false
				}
				seen[cell] = true
			}
		}
	}

	// Check for duplicate values in 3x3 subgrids
	for startRow := 0; startRow < 9; startRow += 3 {
		for startCol := 0; startCol < 9; startCol += 3 {
			seen := make(map[byte]bool)
			for i := 0; i < 3; i++ {
				for j := 0; j < 3; j++ {
					cell := board[startRow+i][startCol+j]
					if cell != '.' {
						if cell < '1' || cell > '9' || seen[cell] {
							return false
						}
						seen[cell] = true
					}
				}
			}
		}
	}

	return true
}
