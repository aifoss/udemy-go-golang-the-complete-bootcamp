package main

import "fmt"

/**
 * Created by sofia on 2019-12-24.
 */

/**
 * Source: Udemy Go (Golang) The Complete Bootcamp Course
 */

func main() {
	var n int

	n = n + 1
	fmt.Println(n)

	n += 2
	fmt.Println(n)

	n++
	fmt.Println(n)

	n = 10

	n = n - 1
	fmt.Println(n)

	n -= 2
	fmt.Println(n)

	n--
	fmt.Println(n)

	var counter int

	counter++
	counter++
	counter++
	counter--

	fmt.Printf("There are %d lines\n", counter)
}
