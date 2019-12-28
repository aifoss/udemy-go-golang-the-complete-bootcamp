package main

import "fmt"

/**
 * Created by sofia on 2019-12-23.
 */

/**
 * Source: Udemy Go (Golang) The Complete Bootcamp Course
 */

func main() {
	var (
		perimeter     int
		width, height = 5, 6
	)

	perimeter = 2 * (width + height)

	fmt.Println("perimeter =", perimeter)
}
