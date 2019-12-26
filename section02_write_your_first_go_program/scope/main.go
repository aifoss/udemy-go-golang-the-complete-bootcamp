package main

// file scope
import "fmt"

/**
 * Created by sofia on 2019-12-22.
 */

/**
 * Source: Udemy Go (Golang) The Complete Bootcamp Course
 */

// package scope
const ok = true

var me = 10

// block scope
func nested() {
	var me = 5
	fmt.Println("inside nested: ", me)
}

// block scope
func main() {
	var hello = "Hello"
	fmt.Println(hello, ok)

	nested()

	fmt.Println("inside main: ", me)

	hey()

	bye()
}
