package main

import "fmt"

/**
 * Created by sofia on 2019-12-24.
 */

/**
 * Source: Udemy Go (Golang) The Complete Bootcamp Course
 */

func main() {
	width, height := 5., 12.

	area := width * height

	fmt.Printf("area=%g\n", area)

	area = area + 10
	area = area - 10
	area = area * 2
	area = area / 2

	fmt.Printf("area=%g\n", area)

	area += 10
	area -= 10
	area *= 2
	area /= 2

	fmt.Printf("area=%g\n", area)
}
