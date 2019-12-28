package main

import "fmt"

/**
 * Created by sofia on 2019-12-26.
 */

/**
 * Source: Udemy Go (Golang) The Complete Bootcamp Course
 */

func main() {
	var letter byte = 'A'
	fmt.Println("English letter:", letter)

	var unicode rune = 'Ç'
	fmt.Println("Non-English letter:", unicode)

	var year uint16 = 2040
	fmt.Println("Year in 4 digits:", year)

	var month uint8 = 6
	fmt.Println("Month in 2 digits:", month)

	var lightSpeed int64 = 670616629
	fmt.Println("Spee of light:", lightSpeed)

	var angle float32 = 35.8
	fmt.Println("Angle of circle:", angle)

	var pi float64 = 3.141592653589793
	fmt.Println("PI number:", pi)
}
