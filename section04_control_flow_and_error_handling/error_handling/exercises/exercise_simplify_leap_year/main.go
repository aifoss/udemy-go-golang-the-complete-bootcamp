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

func main() {
	args := os.Args

	if len(args) != 2 {
		fmt.Println("Give me a year number")
		return
	}

	year, err := strconv.Atoi(args[1])

	if err != nil {
		fmt.Printf("%q is not a valid year\n", args[1])
		return
	}

	// if year is divisible by 400 then is_leap_year
	// else if year is divisible by 100 then not_leap_year
	// else if year is divisible by 4 then is_leap_year
	// else not_leap_year

	if (year%400 == 0 || year%4 == 0) && year%100 != 0 {
		fmt.Printf("%d is a leap year\n", year)
	} else {
		fmt.Printf("%d is not a leap year\n", year)
	}
}
