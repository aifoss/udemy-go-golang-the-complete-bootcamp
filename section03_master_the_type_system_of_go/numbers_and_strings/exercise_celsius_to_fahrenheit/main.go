package main

import (
	"fmt"
	"os"
	"strconv"
)

/**
 * Created by sofia on 2019-12-24.
 */

/**
 * Source: Udemy Go (Golang) The Complete Bootcamp Course
 */

func main() {
	arg := os.Args[1]

	celsius, _ := strconv.ParseFloat(arg, 64)

	// 20C * 1.8 + 32 = 68F
	// F = (9C + 160) / 5

	fahrenheit := (9*celsius + 160) / 5

	fmt.Printf("%g Celsius is %g Fahrenheight\n", celsius, fahrenheit)
}
