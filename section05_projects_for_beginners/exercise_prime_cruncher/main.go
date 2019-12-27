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
// EXERCISE: Crunch the primes
//
//  1. Get numbers from the command-line.
//  2. `for range` over them.
//  4. Convert each one to an int.
//  5. If one of the numbers isn't an int, just skip it.
//  6. Print the ones that are only the prime numbers.
//
// RESTRICTION
//  The user can run the program with any number of
//  arguments.
//
// HINT
//  Find whether a number is prime using this algorithm:
//  https://stackoverflow.com/a/1801446/115363
//
// EXPECTED OUTPUT
//  go run main.go 2 4 6 7 a 9 c 11 x 12 13
//    2 7 11 13
//
//  go run main.go 1 2 3 5 7 A B C
//    2 3 5 7
// ---------------------------------------------------------

const maxInputSize = 100

func isPrime(x int) bool {
	if x < 2 {
		return false
	}
	if x%2 == 0 {
		return x == 2
	}

	for i := 3; i*i <= x; i += 2 {
		if x%i == 0 {
			return false
		}
	}

	return true
}

func main() {
	args := os.Args[1:]

	for i := range args {
		n, err := strconv.Atoi(args[i])

		if err != nil {
			continue
		}

		if isPrime(n) {
			fmt.Printf("%d ", n)
		}
	}

	fmt.Println()
}
