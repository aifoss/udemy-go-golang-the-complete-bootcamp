package main

import "fmt"

/**
 * Created by sofia on 2019-12-24.
 */

/**
 * Source: Udemy Go (Golang) The Complete Bootcamp Course
 */

func main() {
	var (
		myAge   = 30
		yourAge = 35
		average float64
	)

	average = float64(myAge+yourAge) / 2

	fmt.Println(average)

	ratio := 1.0 / 10

	fmt.Printf("%.60f\n", ratio)

	fmt.Println("sum:", 3+2)
	fmt.Println("sum:", 2+3.5)
	fmt.Println("diff:", 3-1)
	fmt.Println("diff:", 3-0.5)
	fmt.Println("prod:", 4*5)
	fmt.Println("prod:", 4*5.5)
	fmt.Println("quot:", 8/3)
	fmt.Println("quot:", 8/3.0)
	fmt.Println("rem:", 5%2)
	fmt.Println("rem:", int(5.0)%2)

	var rate float64 = 3 / 2
	fmt.Println(rate)

	rate = float64(int(3) / int(2))
	fmt.Println(rate)

	rate = float64(3) / 2
	fmt.Println(rate)

	rate = 3.0 / 2
	fmt.Println(rate)
}
