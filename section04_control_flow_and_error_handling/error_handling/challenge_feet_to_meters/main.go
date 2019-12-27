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

const feetToMeters = 0.3048

func main() {
	args := os.Args
	argCount := len(args) - 1

	if argCount != 1 {
		fmt.Println("Provide a number")
		return
	}

	input := args[1]
	feet, err := strconv.ParseInt(input, 10, 64)

	if err != nil {
		fmt.Printf("%q is not a number\n", input)
		return
	}

	meters := float64(feet) * feetToMeters

	fmt.Printf("Parsed %q to %d\n", input, feet)
	fmt.Printf("Converted %d feet to %g meters\n", feet, meters)
}
