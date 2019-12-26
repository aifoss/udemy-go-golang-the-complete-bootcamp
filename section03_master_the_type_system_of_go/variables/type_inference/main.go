package main

import "fmt"

/**
 * Created by sofia on 2019-12-23.
 */

/**
 * Source: Udemy Go (Golang) The Complete Bootcamp Course
 */

// no short declaration for package-level variables
var packageScopeVar bool = true
var packageScopeVar2 = true

func main() {
	// normal declaration
	var name string = "Carl"
	var isScientist bool = true
	var age int = 62
	var degree float64 = 5.

	fmt.Println(name, isScientist, age, degree)

	// normal declaration 2
	var name2 = "Carl"
	var isScientist2 = true
	var age2 = 62
	var degree2 = 5.

	fmt.Println(name2, isScientist2, age2, degree2)

	// short declaration
	name3 := "Carl"
	isScientist3 := true
	age3 := 62
	degree3 := 5.

	fmt.Println(name3, isScientist3, age3, degree3)

	safe, speed := true, 50

	fmt.Println(safe, speed)
}
