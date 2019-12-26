package main

import (
	"fmt"
	"os"
)

/**
 * Created by sofia on 2019-12-24.
 */

/**
 * Source: Udemy Go (Golang) The Complete Bootcamp Course
 */

// go build -o greeter
// ./greeter sofia
func main() {
	name := os.Args[1]

	fmt.Println("Hello great", name, "!")

	name, age := "gandalf", 2019

	fmt.Println("My name is", name)
	fmt.Println("My age is", age)
	fmt.Println("BTW, you shall pass!")
}
