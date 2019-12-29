package main

import "fmt"

/**
 * Created by sofia on 2019-12-28.
 */

/**
 * Source: Udemy Go (Golang) The Complete Bootcamp Course
 */

func main() {
	type person struct {
		firstname, lastname string
		age                 int
	}

	var picasso person

	picasso.firstname = "Pablo"
	picasso.lastname = "Picasso"
	picasso.age = 91

	freud := person{firstname: "Sigmund", lastname: "Freud", age: 83}

	smith := person{
		"John",
		"Smith",
		30,
	}

	fmt.Printf("Picasso: %+v\n", picasso)
	fmt.Printf("Freud: %+v\n", freud)
	fmt.Printf("Smith: %+v\n", smith)
}
