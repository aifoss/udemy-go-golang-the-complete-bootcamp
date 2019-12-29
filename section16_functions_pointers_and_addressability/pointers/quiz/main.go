package main

import "fmt"

/**
 * Created by sofia on 2019-12-28.
 */

/**
 * Source: Udemy Go (Golang) The Complete Bootcamp Course
 */

func main() {
	type computer struct {
		brand string
	}

	var a, b *computer
	fmt.Print(a == b)

	a = &computer{"Apple"}
	b = &computer{"Apple"}
	fmt.Print(" ", a == b, " ", *a == *b, "\n")
}
