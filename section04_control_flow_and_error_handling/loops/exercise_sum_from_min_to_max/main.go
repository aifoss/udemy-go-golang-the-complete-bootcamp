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
// EXERCISE: Sum up to N
//
//  1. Get two numbers from the command-line: min and max
//  2. Convert them to integers (using Atoi)
//  3. By using a loop, sum the numbers between min and max
//
// RESTRICTIONS
//  1. Be sure to handle the errors. So, if a user doesn't
//     pass enough arguments or she passes non-numerics,
//     then warn the user and exit from the program.
//
//  2. Also, check that, min < max.
//
// HINT
//  For converting the numbers, you can use `strconv.Atoi`.
//
// EXPECTED OUTPUT
//  Let's suppose that the user runs it like this:
//
//    go run main.go 1 10
//
//  Then it should print:
//
//    1 + 2 + 3 + 4 + 5 + 6 + 7 + 8 + 9 + 10 = 55
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
		sum += i

		if i < max {
			fmt.Printf("%d + ", i)
		} else {
			fmt.Printf("%d = ", i)
		}
	}

	fmt.Println(sum)
}
