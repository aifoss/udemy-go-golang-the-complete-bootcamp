package main

import "fmt"

/**
 * Created by sofia on 2019-12-23.
 */

/**
 * Source: Udemy Go (Golang) The Complete Bootcamp Course
 */

func main() {
	// at least one of the variables in redeclaration should be a new variable

	var safe bool
	safe, speed := true, 50

	fmt.Println(safe, speed)

	name := "Nikola"
	name, age := "Marie", 60

	fmt.Println(name, age)

	name = "Albert"
	birth := 1879

	fmt.Println(name, birth)
}
