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
// EXERCISE: Compare the slices
//
//  1. Split the namesA string and get a slice
//
//  2. Sort all the slices
//
//  3. Compare whether the slices are equal or not
//
//
// EXPECTED OUTPUT
//
//   They are equal.
//
//
// HINTS
//
//   1. strings.Split function splits a string and
//      returns a string slice
//
//   2. Comparing slices: First check whether their length
//      are the same or not; only then compare them.
//
// ---------------------------------------------------------

func main() {
	namesA := strings.Split("Da Vinci, Wozniak, Carmack", ", ")
	namesB := []string{"Wozniak", "Da Vinci", "Carmack"}

	sort.Strings(namesA)
	sort.Strings(namesB)

	var equal = true

	for i := range namesA {
		if namesA[i] != namesB[i] {
			equal = false
			break
		}
	}

	switch equal {
	case true:
		fmt.Println("They are equal.")
	case false:
		fmt.Println("They are NOT equal.")
	}
}
