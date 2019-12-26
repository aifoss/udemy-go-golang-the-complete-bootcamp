package main

import "fmt"

/**
 * Created by sofia on 2019-12-23.
 */

/**
 * Source: Udemy Go (Golang) The Complete Bootcamp Course
 */

func main() {
	var speed int
	fmt.Println(speed)

	speed = 100
	fmt.Println(speed)

	speed -= 25
	fmt.Println(speed)

	var (
		name   string
		age    int
		famous bool
	)

	name = "Newton"
	age = 84
	famous = true

	fmt.Println(name, age, famous)

	var prevName string
	prevName = name

	name = "Einstein"

	fmt.Println("previous name: ", prevName)
	fmt.Println("current name: ", name)
}
