package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

/**
 * Created by sofia on 2019-12-26.
 */

/**
 * Source: Udemy Go (Golang) The Complete Bootcamp Course
 */

const usage = `
Feet To Meters
--------------
This program converts feet to meters.

Usage:
feet [feetToConvert]
`

func main() {
	if len(os.Args) < 2 {
		fmt.Println(strings.TrimSpace(usage))
		return
	}

	arg := os.Args[1]

	feet, _ := strconv.ParseFloat(arg, 64)

	meters := feet * 0.3048

	fmt.Printf("%g feet is %g meters\n", feet, meters)
}
