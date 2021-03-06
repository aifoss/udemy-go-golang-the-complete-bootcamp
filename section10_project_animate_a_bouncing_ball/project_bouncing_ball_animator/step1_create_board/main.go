package main

import "fmt"

/**
 * Created by sofia on 2019-12-27.
 */

/**
 * Source: Udemy Go (Golang) The Complete Bootcamp Course
 */

func main() {
	const (
		width  = 50
		height = 10

		cellEmpty = ' '
		cellBall  = '⚾'
	)

	board := make([][]bool, width)
	for row := range board {
		board[row] = make([]bool, height)
	}

	board[0][0] = true

	var cell rune

	for y := range board[0] {
		for x := range board {
			cell = cellEmpty
			if board[x][y] {
				cell = cellBall
			}
			fmt.Print(string(cell), " ")
		}
		fmt.Println()
	}
}
