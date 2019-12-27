package main

import (
	"fmt"
)

/**
 * Created by sofia on 2019-12-26.
 */

/**
 * Source: Udemy Go (Golang) The Complete Bootcamp Course
 */

func main() {
	i := 10

	// switch true {
	switch {
	case i > 0:
		fmt.Println("positive")
	case i < 0:
		fmt.Println("negative")
	default:
		fmt.Println("zero")
	}
}
