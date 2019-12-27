package main

import (
	"fmt"
	"os"
)

/**
 * Created by sofia on 2019-12-26.
 */

/**
 * Source: Udemy Go (Golang) The Complete Bootcamp Course
 */

func main() {
	city := os.Args[1]

	switch city {
	case "Paris":
		fmt.Println("France")
	case "Tokyo":
		fmt.Println("Japan")
	default:
		fmt.Println("Where?")
	}
}
