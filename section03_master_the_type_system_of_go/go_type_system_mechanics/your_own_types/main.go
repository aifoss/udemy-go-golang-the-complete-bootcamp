package main

import "fmt"

/**
 * Created by sofia on 2019-12-25.
 */

/**
 * Source: Udemy Go (Golang) The Complete Bootcamp Course
 */

type gram float64
type ounce float64

func main() {
	var g gram = 1000
	var o ounce
	o = ounce(g) * 0.035274

	fmt.Printf("%g grams is %.2f ounces\n", g, o)
}
