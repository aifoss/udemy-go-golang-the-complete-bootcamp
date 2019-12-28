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

	var nums []int
	s.Show("no backing array", nums)

	nums = append(nums, 1, 3)
	s.Show("after appending 1, 3", nums)

	nums = append(nums, 4)
	s.Show("after appending 4", nums)

	nums = append(nums, 5)
	s.Show("after appending 5", nums)

	nums = append(nums, nums[2:]...)
	s.Show("after appending last 2 values", nums)

	nums = append(nums[:2], 7, 9)
	s.Show("appending 7, 9 after first 2 values", nums)

	nums = nums[:6]
	s.Show("after extending to 6", nums)

	nums = nums[:8]
	s.Show("after extending to 8", nums)
}
