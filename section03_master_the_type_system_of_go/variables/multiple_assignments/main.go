package main

import (
	"fmt"
	"time"
)

/**
 * Created by sofia on 2019-12-23.
 */

/**
 * Source: Udemy Go (Golang) The Complete Bootcamp Course
 */

func main() {
	var (
		day int
		now time.Time
	)

	day, now = 1, time.Now()

	fmt.Println(day, now)

	// use multiple assignments to swap values of variables
	var (
		speed     = 100
		prevSpeed = 50
	)

	fmt.Println(prevSpeed, speed)

	prevSpeed, speed = speed, prevSpeed

	fmt.Println(prevSpeed, speed)
}
