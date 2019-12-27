package main

import "fmt"

/**
 * Created by sofia on 2019-12-26.
 */

/**
 * Source: Udemy Go (Golang) The Complete Bootcamp Course
 */

func main() {
	var sum int

	for i := 1; i <= 1000; i++ {
		sum += i
	}

	fmt.Println(sum)
}
