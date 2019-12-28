package main

import (
	"fmt"
	"sort"
	"strings"
)

/**
 * Created by sofia on 2019-12-27.
 */

/**
 * Source: Udemy Go (Golang) The Complete Bootcamp Course
 */

// ---------------------------------------------------------
// EXERCISE: Fix the problems
//
//  1. Uncomment the code
//
//  2. Fix the problems
//
//  3. BONUS: Simplify the code
//
//
// EXPECTED OUTPUT
//   "Einstein and Shepard and Tesla"
//   ["Fire" "Kafka's Revenge" "Stay Golden"]
//   [1 2 3 5 6 7 8 9]
// ---------------------------------------------------------

func main() {
	var names []string
	names = []string{
		"Einstein",
		"Shepard",
		"Tesla",
	}

	//-----------------------------------
	var books []string = []string{
		"Stay Golden",
		"Fire",
		"Kafka's Revenge",
	}

	sort.Strings(books)

	//-----------------------------------
	// this time, do not change the nums array to a slice
	nums := [...]int{5, 1, 7, 3, 8, 2, 6, 9}

	// use the slicing expression to change the nums array to a slice below
	sort.Ints(nums[:])

	//-----------------------------------
	// Here: Use the strings.Join function to join the names
	//      (see the expected output)
	nameStr := strings.Join(names, " and ")
	fmt.Printf("%q\n", nameStr)

	fmt.Printf("%q\n", books)
	fmt.Printf("%d\n", nums)
}
