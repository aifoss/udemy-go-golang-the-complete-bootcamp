package main

import (
	"fmt"
	"os"
	"strconv"
)

/**
 * Created by sofia on 2019-12-27.
 */

/**
 * Source: Udemy Go (Golang) The Complete Bootcamp Course
 */

// ---------------------------------------------------------
// EXERCISE: Number Sorter
//
//  Your goal is to sort the given numbers from the command-line.
//
//  1. Get the numbers from the command-line.
//
//  2. Create an array and assign the given numbers to that array.
//
//  3. Sort the given numbers and print them.
//
// RESTRICTION
//   + Maximum 5 numbers can be provided
//   + If one of the arguments are not a valid number, skip it
//
// HINTS
//  + You can use the bubble-sort algorithm to sort the numbers.
//    Please watch this: https://youtu.be/nmhjrI-aW5o?t=7
//
//  + When swapping the elements, do not check for the last element.
//
//    Or, you will receive this error:
//    "panic: runtime error: index out of range"
//
// EXPECTED OUTPUT
//   go run main.go
//     Please give me up to 5 numbers.
//
//   go run main.go 6 5 4 3 2 1
//     Sorry. Go arrays are fixed. So, for now, I'm only supporting sorting 5 numbers...
//
//   go run main.go 5 4 3 2 1
//     [1 2 3 4 5]
//
//   go run main.go 5 4 a c 1
//     [0 0 1 4 5]
// ---------------------------------------------------------

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		fmt.Println("Please give me up to 5 numbers.")
		return
	} else if len(args) > 5 {
		fmt.Println("Sorry. Go arrays are fixed. So, for now, I'm only supporting sorting 5 numbers.")
		return
	}

	var array [5]int

	for i, v := range args {
		number, err := strconv.Atoi(v)

		if err != nil {
			continue
		}

		array[i] = number
	}

	// n := len(array)
	// for i := 0; i < n-1; i++ {
	// 	for j := 0; j < n-1-i; j++ {
	// 		if array[j] > array[j+1] {
	// 			// fmt.Printf("swapping %d and %d\n", array[j], array[j+1])
	// 			array[j], array[j+1] = array[j+1], array[j]
	// 		}
	// 	}
	// }

	for range array {
		for i, v := range array {
			if i < len(array)-1 && v > array[i+1] {
				array[i], array[i+1] = array[i+1], array[i]
			}
		}
	}

	fmt.Printf("%v\n", array)
}
