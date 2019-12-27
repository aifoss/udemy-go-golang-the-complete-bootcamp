package main

import (
	"fmt"
	"strings"
)

/**
 * Created by sofia on 2019-12-26.
 */

/**
 * Source: Udemy Go (Golang) The Complete Bootcamp Course
 */

func main() {
	fmt.Println(strings.Repeat("-", 25))

	speed := 10
	speedB := 150.5
	faster := speedB > 100

	fmt.Println("is the other car going faster?", faster)

	faster = speedB > float64(speed)

	fmt.Println("is the other car going faster?", faster)
}
