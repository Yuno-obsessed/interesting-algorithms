package main

import (
	"fmt"
)

var gridSize = 9

var grid = [9][9]int{
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

func isPresentInCol(col, num int) bool {
	for row := 0; row < gridSize; row++ {
		if grid[row][col] == num {
			return true
		}
	}
	return false
}

func isPresentInRow(row, num int) bool {
	for col := 0; col < gridSize; col++ {
		if grid[row][col] == num {
			return true
		}
	}
	return false
}

func isPresentInBox(boxStartRow, boxStartCol, num int) bool {
	for row := 0; row < 3; row++ {
		for col := 0; col < 3; col++ {
			if grid[row+boxStartRow][col+boxStartCol] == num {
				return true
			}
		}
	}
	return false
}

func sudokuGrid() {
	for row := 0; row < gridSize; row++ {
		for col := 0; col < gridSize; col++ {
			if col == 3 || col == 6 {
				fmt.Print(" | ")
			}
			fmt.Print(grid[row][col])
			fmt.Print(" ")
		}

		if row == 2 || row == 5 {
			fmt.Println()
			for i := 0; i < gridSize; i++ {
				fmt.Print("---")
			}
		}
		fmt.Println()
	}
}

func findEmptyPlace(row int, col int) (bool, int, int) {
	for row := 0; row < gridSize; row++ {
		for col := 0; col < gridSize; col++ {
			if grid[row][col] == 0 {
				return true, row, col
			}
		}
	}
	return false, 0, 0
}

func isValidPlace(row, col, num int) bool {
	return !isPresentInRow(row, num) && !isPresentInCol(col, num) && !isPresentInBox(row-row%3, col-col%3, num)
}

func solveSudoku() bool {
	var row, col int
	ok, row, col := findEmptyPlace(row, col)
	if !ok {
		return true
	}
	for num := 1; num <= 9; num++ {
		if isValidPlace(row, col, num) {
			grid[row][col] = num
			if solveSudoku() {
				return true
			}
			grid[row][col] = 0
		}
	}
	return false
}

func main() {
	if solveSudoku() {
		sudokuGrid()
	} else {
		fmt.Println("No solution exists")
	}
}
