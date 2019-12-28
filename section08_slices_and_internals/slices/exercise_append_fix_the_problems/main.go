package main

import "fmt"

/**
 * Created by sofia on 2019-12-27.
 */

/**
 * Source: Udemy Go (Golang) The Complete Bootcamp Course
 */

// ---------------------------------------------------------
// EXERCISE: Append #3 — Fix the problems
//
//  Fix the problems in the code below.
//
// BONUS
//
//  Simplify the code.
//
// EXPECTED OUTPUT
//
//  toppings: [black olives green peppers onions extra cheese]
//
// ---------------------------------------------------------

func main() {
	toppings := []string{"black olives", "green peppers"}

	var pizza []string

	pizza = append(pizza, toppings...)

	pizza = append(toppings, "onions")
	pizza = append(pizza, "extra cheese")

	fmt.Printf("pizza       : %s\n", pizza)
}
