package main

import "fmt"

/**
 * Created by sofia on 2019-12-27.
 */

/**
 * Source: Udemy Go (Golang) The Complete Bootcamp Course
 */

func main() {
	var (
		blue = [3]int{6, 9, 3}
		red  = [...]int{6, 9, 3}
	)

	fmt.Printf("blue bookcase: %v\n", blue)
	fmt.Printf("red bookcase: %v\n", red)

	fmt.Println("Are they equal?", blue == red)
}
