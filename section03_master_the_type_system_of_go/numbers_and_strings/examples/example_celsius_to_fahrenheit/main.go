package main

import "fmt"

/**
 * Created by sofia on 2019-12-24.
 */

/**
 * Source: Udemy Go (Golang) The Complete Bootcamp Course
 */

func main() {
	celsius := 35.
	fahrenheit := (9*celsius + 160) / 5

	fmt.Printf("%g C is %g F\n", celsius, fahrenheit)
}
