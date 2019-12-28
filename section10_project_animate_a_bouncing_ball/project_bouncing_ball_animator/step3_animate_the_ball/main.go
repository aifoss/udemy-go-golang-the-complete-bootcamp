package main

import (
	"fmt"
	"time"

	"github.com/inancgumus/screen"
)

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

		maxFrames = 1200
		sleep     = time.Second / 20

		cellEmpty = ' '
		cellBall  = '⚾'
	)

	var (
		px, py int
		vx, vy = 1, 1
		cell   rune
	)

	board := make([][]bool, width)
	for row := range board {
		board[row] = make([]bool, height)
	}

	screen.Clear()

	for i := 0; i < maxFrames; i++ {
		px += vx
		py += vy

		if px == 0 || px >= width-1 {
			vx *= -1
		}
		if py == 0 || py >= height-1 {
			vy *= -1
		}

		for y := range board[0] {
			for x := range board {
				board[x][y] = false
			}
		}

		board[px][py] = true

		buf := make([]rune, 0, width*height)
		buf = buf[:0]

		for y := range board[0] {
			for x := range board {
				cell = cellEmpty
				if board[x][y] {
					cell = cellBall
				}
				buf = append(buf, cell, ' ')
			}
			buf = append(buf, '\n')
		}

		screen.MoveTopLeft()
		fmt.Println(string(buf))

		time.Sleep(sleep)
	}
}
