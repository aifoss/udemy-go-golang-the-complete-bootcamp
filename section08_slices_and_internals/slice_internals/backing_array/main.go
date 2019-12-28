package main

import (
	"fmt"
	"sort"

	s "github.com/inancgumus/prettyslice"
)

/**
 * Created by sofia on 2019-12-27.
 */

/**
 * Source: Udemy Go (Golang) The Complete Bootcamp Course
 */

func main() {
	grades := []float64{40, 10, 20, 50, 60, 70}

	var newGrades []float64
	newGrades = append([]float64(nil), grades...)

	front := newGrades[:3]
	front2 := front[:3]
	front3 := front

	sort.Float64s(front)

	s.PrintBacking = true
	s.MaxPerLine = 7
	s.Show("grades", grades[:])
	s.Show("newGrades", newGrades)
	s.Show("front", front)
	s.Show("front2", front2)
	s.Show("front3", front3)

	arr := [...]int{9, 7, 5, 3, 1}
	nums := arr[2:]
	nums2 := nums[1:]

	arr[2]++
	nums[1] -= arr[4] - 4
	nums2[1] += 5

	fmt.Println(nums)
}
