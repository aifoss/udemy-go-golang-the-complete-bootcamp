package main

import "fmt"

/**
 * Created by sofia on 2019-12-27.
 */

/**
 * Source: Udemy Go (Golang) The Complete Bootcamp Course
 */

func main() {
	rates := [3]float64{
		1: 2.5,
		0: 0.5,
		2: 1.5,
	}

	fmt.Printf("%v\n", rates)

	rates = [3]float64{
		2: 1.5,
	}

	fmt.Printf("%v\n", rates)

	rates2 := [...]float64{
		5: 1.5,
	}

	fmt.Printf("%v\n", rates2)

	rates3 := [...]float64{
		5: 1.5,
		2.5,
		0: 0.5,
	}

	fmt.Printf("%v\n", rates3)
}
