package main

import (
	"fmt"
	"strconv"
	"strings"
)

/**
 * Created by sofia on 2019-12-27.
 */

/**
 * Source: Udemy Go (Golang) The Complete Bootcamp Course
 */

// ---------------------------------------------------------
// EXERCISE: Slice the numbers
//
//   We've a string that contains even and odd numbers.
//
//   1. Convert the string to an []int
//
//   2. Print the slice
//
//   3. Slice it for the even numbers and print it (assign it to a new slice variable)
//
//   4. Slice it for the odd numbers and print it (assign it to a new slice variable)
//
//   5. Slice it for the two numbers at the middle
//
//   6. Slice it for the first two numbers
//
//   7. Slice it for the last two numbers (use the len function)
//
//   8. Slice the evens slice for the last one number
//
//   9. Slice the odds slice for the last two numbers
//
//
// EXPECTED OUTPUT
//   go run main.go
//
//     nums        : [2 4 6 1 3 5]
//     evens       : [2 4 6]
//     odds        : [1 3 5]
//     middle      : [6 1]
//     first 2     : [2 4]
//     last 2      : [3 5]
//     evens last 1: [6]
//     odds last 2 : [3 5]
//
//
// NOTE
//
//  You can also use my prettyslice package for printing the slices.
//
//
// HINT
//
//  Find a function in the strings package for splitting the string into []string
//
// ---------------------------------------------------------

func main() {
	// uncomment the declaration below
	data := "2 4 6 1 3 5"
	strs := strings.Split(data, " ")

	var nums []int

	for _, s := range strs {
		num, _ := strconv.Atoi(s)
		nums = append(nums, num)
	}
	var (
		evens      []int
		odds       []int
		middle     []int
		first2     []int
		last2      []int
		evensLast1 []int
		oddsLast2  []int
	)

	for _, v := range nums {
		if v%2 == 0 {
			evens = append(evens, v)
		} else {
			odds = append(odds, v)
		}
	}

	l := len(nums)

	mid1, mid2 := l/2-1, l/2

	middle = nums[mid1 : mid2+1]

	first2 = nums[:2]
	last2 = nums[l-2:]
	evensLast1 = evens[len(evens)-1:]
	oddsLast2 = odds[len(odds)-2:]

	fmt.Printf("nums        : %v\n", nums)
	fmt.Printf("evens       : %v\n", evens)
	fmt.Printf("odds        : %v\n", odds)
	fmt.Printf("middle      : %v\n", middle)
	fmt.Printf("first 2     : %v\n", first2)
	fmt.Printf("last 2      : %v\n", last2)
	fmt.Printf("evens last 1: %v\n", evensLast1)
	fmt.Printf("odds last 2 : %v\n", oddsLast2)
}
