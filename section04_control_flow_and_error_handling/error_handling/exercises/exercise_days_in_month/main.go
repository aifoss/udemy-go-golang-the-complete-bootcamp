package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

/**
 * Created by sofia on 2019-12-26.
 */

/**
 * Source: Udemy Go (Golang) The Complete Bootcamp Course
 */

func isLeapYear(year int) bool {
	return (year%400 == 0 || year%4 == 0) && year%100 != 0
}

func main() {
	args := os.Args

	if len(args) != 2 {
		fmt.Println("Give me a month name")
		return
	}

	year := time.Now().Year()
	month := strings.ToLower(args[1])
	days := 28

	if month == "january" ||
		month == "march" ||
		month == "may" ||
		month == "july" ||
		month == "august" ||
		month == "october" ||
		month == "december" {
		days = 31
	} else if month == "april" ||
		month == "june" ||
		month == "september" ||
		month == "november" {
		days = 30
	} else if month == "february" {
		if isLeapYear(year) {
			days = 29
		}
	} else {
		fmt.Printf("%q is not a valid month name\n", month)
	}

	fmt.Printf("%q has %d days\n", month, days)
}
