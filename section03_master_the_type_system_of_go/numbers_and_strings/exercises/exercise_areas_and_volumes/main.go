package main

import (
	"fmt"
	"math"
)

/**
 * Created by sofia on 2019-12-24.
 */

/**
 * Source: Udemy Go (Golang) The Complete Bootcamp Course
 */

func main() {
	// circle area
	var (
		radius = 10.
		area   float64
	)

	area = math.Pi * radius * radius

	fmt.Printf("radius: %g -> circle area: %.2f\n", radius, area)

	area *= area

	fmt.Printf("radius: %g -> sphere area: %.2f\n", radius, area)

	volume := (float64(4) / float64(3)) * math.Pi * math.Pow(radius, 2) * radius

	fmt.Printf("radius: %g -> sphere volume: %.2f\n", radius, volume)
}
