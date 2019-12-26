package main

import "fmt"

/**
 * Created by sofia on 2019-12-24.
 */

/**
 * Source: Udemy Go (Golang) The Complete Bootcamp Course
 */

func main() {
	var brand = "Blah"

	fmt.Printf("%q\n", brand)
	fmt.Printf("%s\n", brand)

	var ops = 2350
	var ok = 533
	var fail = 433

	fmt.Println(
		"total:", ops, "success:", ok, "/", fail,
	)
	fmt.Printf(
		"total: %d success: %d / %d\n",
		ops, ok, fail,
	)
}
