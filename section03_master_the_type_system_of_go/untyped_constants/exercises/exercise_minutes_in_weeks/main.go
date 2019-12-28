package main

import "fmt"

/**
 * Created by sofia on 2019-12-26.
 */

/**
 * Source: Udemy Go (Golang) The Complete Bootcamp Course
 */

func main() {
	const minsPerDay = 24 * 60
	const weekDays = 7

	fmt.Printf("Number of minutes in 2 weeks: %d\n", minsPerDay*weekDays*2)
}
