package main

import "fmt"

/**
 * Created by sofia on 2019-12-26.
 */

/**
 * Source: Udemy Go (Golang) The Complete Bootcamp Course
 */

func main() {
	var i int

loop:
	if i < 3 {
		fmt.Println("looping")
		i++
		goto loop
	}

	fmt.Println("done")
}
