package main

import "fmt"

/**
 * Created by sofia on 2019-12-26.
 */

/**
 * Source: Udemy Go (Golang) The Complete Bootcamp Course
 */

func main() {
	const min int = 1
	const pi float64 = 3.14
	const version string = "2.0.1"
	const debug bool = true

	fmt.Println(min, pi, version, debug)

	const min2 = 1
	const pi2 = 3.14
	const version2 = "2.0.1"
	const debug2 = true

	fmt.Println(min2, pi2, version2, debug2)

	const min3 = 1 + 1
	const pi3 = 3.14 * min3
	const version3 = "2.0.1" + "-beta"
	const debug3 = !true

	fmt.Println(min3, pi3, version3, debug3)
}
