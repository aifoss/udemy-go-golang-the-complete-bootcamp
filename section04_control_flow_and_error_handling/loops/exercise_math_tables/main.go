package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

/**
 * Created by sofia on 2019-12-26.
 */

/**
 * Source: Udemy Go (Golang) The Complete Bootcamp Course
 */

// ---------------------------------------------------------
// EXERCISE: Math Tables
//
//  Create division, addition and subtraction tables
//
//  1. Get the math operation and
//     the size of the table from the user
//
//  2. Print the table accordingly
//
//  3. Correctly handle the division by zero error
//
//
// BONUS #1
//
//  Use strings.IndexAny function to detect
//    the valid operations.
//
//  Search on Google for: golang pkg strings IndexAny
//
// BONUS #2
//
//  Add remainder operation as well (remainder table using %).
//
//
// EXPECTED OUTPUT
//
//  go run main.go
//    Usage: [op=*/+-] [size]
//
//  go run main.go "*"
//    Size is missing
//    Usage: [op=*/+-] [size]
//
//  go run main.go "%" 4
//    Invalid operator.
//    Valid ops one of: */+-
//
//  go run main.go "*" 4
//    *    0    1    2    3    4
//    0    0    0    0    0    0
//    1    0    1    2    3    4
//    2    0    2    4    6    8
//    3    0    3    6    9   12
//    4    0    4    8   12   16
//
//  go run main.go "/" 4
//    /    0    1    2    3    4
//    0    0    0    0    0    0
//    1    0    1    0    0    0
//    2    0    2    1    0    0
//    3    0    3    1    1    0
//    4    0    4    2    1    1
//
//  go run main.go "+" 4
//    +    0    1    2    3    4
//    0    0    1    2    3    4
//    1    1    2    3    4    5
//    2    2    3    4    5    6
//    3    3    4    5    6    7
//    4    4    5    6    7    8
//
//  go run main.go "-" 4
//    -    0    1    2    3    4
//    0    0   -1   -2   -3   -4
//    1    1    0   -1   -2   -3
//    2    2    1    0   -1   -2
//    3    3    2    1    0   -1
//    4    4    3    2    1    0
//
// BONUS:
//
//  go run main.go "%" 4
//    %    0    1    2    3    4
//    0    0    0    0    0    0
//    1    0    0    1    1    1
//    2    0    0    0    2    2
//    3    0    0    1    0    3
//    4    0    0    0    1    0
//
// NOTES:
//
//   When running the program in Windows Git Bash, you might need
//   to escape the characters like this.
//
//   This happens because those characters have special meaning.
//
//   Division:
//     go run main.go // 4
//
//   Multiplication and others:
//   (this is also necessary for non-Windows bashes):
//
//     go run main.go "*" 4
// ---------------------------------------------------------

const usage = "Usage: [op=*/+-] [size]"
const ops = "*/+-"

func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Println("Missing op and size")
		return
	}

	if len(args) < 3 {
		fmt.Println("Missing size")
		return
	}

	op := args[1]

	if strings.IndexAny(op, ops) == -1 {
		fmt.Println("Invalid operator")
		fmt.Println("Valid op is one of: */+-")
		return
	}

	size, err := strconv.Atoi(args[2])

	if err != nil || size <= 0 {
		fmt.Println("Invalid size")
		return
	}

	fmt.Printf("%5s", "X")

	for i := 1; i <= size; i++ {
		fmt.Printf("%5d", i)
	}
	fmt.Println()

	for i := 1; i <= size; i++ {
		fmt.Printf("%5d", i)

		for j := 1; j <= size; j++ {
			val := 0

			switch op {
			case "*":
				val = i * j
			case "/":
				val = i / j
			case "+":
				val = i + j
			case "-":
				val = i - j
			}

			fmt.Printf("%5d", val)
		}

		fmt.Println()
	}

}
