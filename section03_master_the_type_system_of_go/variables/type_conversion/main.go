package main

import "fmt"

/**
 * Created by sofia on 2019-12-23.
 */

/**
 * Source: Udemy Go (Golang) The Complete Bootcamp Course
 */

func main() {
	speed := 100
	force := 2.5

	speed = int(float64(speed) * force)

	fmt.Println(speed)
}
