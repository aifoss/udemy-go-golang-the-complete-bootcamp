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
// EXERCISE: Break Up
//
//  1. Extend the "Only Evens" exercise
//  2. This time, use an infinite loop.
//  3. Break the loop when it reaches to the `max`.
//
// RESTRICTIONS
//  You should use the `break` statement once.
//
// HINT
//  Do not forget incrementing the `i` before the `continue`
//  statement and at the end of the loop.
//
// EXPECTED OUTPUT
//    go run main.go 1 10
//    2 + 4 + 6 + 8 + 10 = 30
// --------------------------------------------------------

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
	i := min

	for {
		if i > max {
			break
		}

		if i%2 != 0 {
			i++
			continue
		}

		sum += i

		if i < max {
			fmt.Printf("%d + ", i)
		} else {
			fmt.Printf("%d = ", i)
		}

		i++
	}

	fmt.Println(sum)

}
