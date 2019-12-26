package main

import "fmt"

/**
 * Created by sofia on 2019-12-26.
 */

/**
 * Source: Udemy Go (Golang) The Complete Bootcamp Course
 */

func main() {
	const min = 1 + 1
	const pi = 3.14 * min

	fmt.Println(min, pi)

	const val = 42

	var i int = val
	var f float64 = val
	var b byte = val
	var j int32 = val
	var r rune = val

	fmt.Println(i, f, b, j, r)
}
