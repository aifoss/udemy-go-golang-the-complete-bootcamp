package main

import "fmt"

/**
 * Created by sofia on 2019-12-25.
 */

/**
 * Source: Udemy Go (Golang) The Complete Bootcamp Course
 */

func main() {
	fmt.Printf("%b\n", 0)
	fmt.Printf("%b\n", 1)
	fmt.Printf("%02b\n", 0)
	fmt.Printf("%02b\n", 1)
	fmt.Printf("%02b\n", 2)
	fmt.Printf("%02b\n", 3)
	fmt.Printf("%02b\n", 4)

	var b byte

	b = 0
	fmt.Printf("%08b = %d\n", b, b)

	b = 255
	fmt.Printf("%08b = %d\n", b, b)
}
