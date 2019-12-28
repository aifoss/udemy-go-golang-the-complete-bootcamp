package main

import (
	"fmt"
	"os"
	"strconv"
)

/**
 * Created by sofia on 2019-12-26.
 */

/**
 * Source: Udemy Go (Golang) The Complete Bootcamp Course
 */

// ---------------------------------------------------------
// EXERCISE: Only Evens
//
//  1. Extend the "Sum up to N" exercise
//  2. Sum only the even numbers
//
// RESTRICTIONS
//  Skip odd numbers using the `continue` statement
//
// EXPECTED OUTPUT
//  Let's suppose that the user runs it like this:
//
//    go run main.go 1 10
//
//  Then it should print:
//
//    2 + 4 + 6 + 8 + 10 = 30
// ---------------------------------------------------------

func main() {
	args := os.Args

	if len(args) < 3 {
		fmt.Println("Give min and max")
		return
	}

	min, err := strconv.Atoi(args[1])

	if err != nil {
		fmt.Printf("Invalid min: %q\n", args[1])
		return
	}

	max, err := strconv.Atoi(args[2])

	if err != nil {
		fmt.Printf("Invalid max: %q\n", args[2])
		return
	}

	if min > max {
		fmt.Println("Min should be <= max")
		return
	}

	sum := 0

	for i := min; i <= max; i++ {
		if i%2 != 0 {
			continue
		}

		sum += i

		if i < max {
			fmt.Printf("%d + ", i)
		} else {
			fmt.Printf("%d = ", i)
		}
	}

	fmt.Println(sum)
}
