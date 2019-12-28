package main

import s "github.com/inancgumus/prettyslice"

/**
 * Created by sofia on 2019-12-27.
 */

/**
 * Source: Udemy Go (Golang) The Complete Bootcamp Course
 */

func main() {
	s.PrintBacking = true

	nums := []int{1, 3, 2, 4}
	odds := append(nums[:2:2], 5, 7)
	evens := append(nums[2:4], 6, 8)

	s.Show("nums", nums)
	s.Show("odds", odds)
	s.Show("evens", evens)

}
