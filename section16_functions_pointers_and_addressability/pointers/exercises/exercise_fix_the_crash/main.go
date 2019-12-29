package main

import "fmt"

/**
 * Created by sofia on 2019-12-28.
 */

/**
 * Source: Udemy Go (Golang) The Complete Bootcamp Course
 */

// ---------------------------------------------------------
// EXERCISE: Fix the crash
//
// EXPECTED OUTPUT
//
//   brand: apple
// ---------------------------------------------------------

type computer struct {
	brand *string
}

func main() {
	c := &computer{}
	change(c, "apple")
	fmt.Printf("brand: %s\n", *c.brand)
}

func change(c *computer, brand string) {
	c.brand = &brand
}
