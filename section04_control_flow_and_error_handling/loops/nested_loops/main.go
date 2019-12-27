package main

import "fmt"

/**
 * Created by sofia on 2019-12-26.
 */

/**
 * Source: Udemy Go (Golang) The Complete Bootcamp Course
 */

const max = 5

func main() {
	fmt.Printf("%5s", "X")

	for i := 1; i <= max; i++ {
		fmt.Printf("%5d", i)
	}
	fmt.Println()

	for i := 1; i <= max; i++ {
		fmt.Printf("%5d", i)

		for j := 1; j <= max; j++ {
			fmt.Printf("%5d", i*j)
		}

		fmt.Println()
	}

}
