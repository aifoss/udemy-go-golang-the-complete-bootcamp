package main

import "fmt"

/**
 * Created by sofia on 2019-12-26.
 */

/**
 * Source: Udemy Go (Golang) The Complete Bootcamp Course
 */

func main() {
	const min, max int = 1, 1000

	fmt.Println(min, max)

	const (
		min2 int = 1
		max2 int = 1
	)

	fmt.Println(min2, max2)

	const (
		min3 int = 1
		max3
	)

	fmt.Println(min3, max3)
}
