package main

import "fmt"

/**
 * Created by sofia on 2019-12-26.
 */

/**
 * Source: Udemy Go (Golang) The Complete Bootcamp Course
 */

func main() {
	i := 42
	fmt.Printf("%T\n", i)

	f := 3.14
	fmt.Printf("%T\n", f)

	b := true
	fmt.Printf("%T\n", b)

	s := "Hello"
	fmt.Printf("%T\n", s)

	r := 'A'
	fmt.Printf("%T\n", r)

	ff := 42 + 3.14
	fmt.Printf("%T\n", ff)
}
