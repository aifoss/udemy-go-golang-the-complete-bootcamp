package main

import (
	"math/rand"
	"time"

	s "github.com/inancgumus/prettyslice"
	"github.com/inancgumus/screen"
)

/**
 * Created by sofia on 2019-12-27.
 */

/**
 * Source: Udemy Go (Golang) The Complete Bootcamp Course
 */

func main() {
	s.PrintBacking = true
	s.MaxPerLine = 30
	s.Width = 150

	var nums []int

	screen.Clear()

	for cap(nums) <= 128 {
		screen.MoveTopLeft()

		s.Show("nums", nums)
		nums = append(nums, rand.Intn(9)+1)

		time.Sleep(time.Second / 4)
	}
}
