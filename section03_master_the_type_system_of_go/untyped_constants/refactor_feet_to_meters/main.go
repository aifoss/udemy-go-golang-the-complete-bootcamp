package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

/**
 * Created by sofia on 2019-12-26.
 */

/**
 * Source: Udemy Go (Golang) The Complete Bootcamp Course
 */

func main() {
	const (
		feetToMeters = 0.3048
		feetToYards  = feetToMeters / 0.9144
	)

	arg := os.Args[1]

	feet, _ := strconv.ParseFloat(arg, 64)

	meters := feet * feetToMeters
	yards := math.Round(feet * feetToYards)

	fmt.Printf("%g feet is %g meters\n", feet, meters)
	fmt.Printf("%g feet is %g yards\n", feet, yards)
}
