package main

import "fmt"

/**
 * Created by sofia on 2019-12-27.
 */

/**
 * Source: Udemy Go (Golang) The Complete Bootcamp Course
 */

// ---------------------------------------------------------
// EXERCISE: Fix the backing array problem
//
//  Ensure that changing the elements of the `mine` slice
//  does not change the elements of the `nums` slice.
//
//
// CURRENT OUTPUT (INCORRECT)
//
//  Mine         : [-50 -100 -150 25 30 50]
//  Original nums: [-50 -100 -150]
//
//
// EXPECTED OUTPUT
//
//  Mine         : [-50 -100 -150]
//  Original nums: [56 89 15]
//
// ---------------------------------------------------------

func main() {
	// DON'T TOUCH THE FOLLOWING CODE
	nums := []int{56, 89, 15, 25, 30, 50}

	// ----------------------------------------
	// ONLY ADD YOUR CODE HERE
	//
	// Ensure that nums slice never changes even though
	// the mine slice changes.

	// mine := nums

	// var mine []int
	// mine = append(mine, nums[:3]...)

	mine := append([]int(nil), nums[:3]...)

	// ----------------------------------------

	// DON'T TOUCH THE FOLLOWING CODE
	//
	// This code changes the elements of the nums
	// slice.
	//
	mine[0], mine[1], mine[2] = -50, -100, -150

	fmt.Println("Mine         :", mine[:3])
	fmt.Println("Original nums:", nums[:3])
}
