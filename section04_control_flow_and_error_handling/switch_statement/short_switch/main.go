package main

import "fmt"

/**
 * Created by sofia on 2019-12-26.
 */

/**
 * Source: Udemy Go (Golang) The Complete Bootcamp Course
 */

func main() {
	// switch i := 10; true {
	switch i := 10; {
	case i > 0:
		fmt.Print("positive")
	case i < 0:
		fmt.Print("negative")
	default:
		fmt.Print("zero")
	}

	fmt.Println()
}
