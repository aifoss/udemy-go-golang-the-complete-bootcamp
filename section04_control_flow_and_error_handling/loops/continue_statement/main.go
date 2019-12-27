package main

import "fmt"

/**
 * Created by sofia on 2019-12-26.
 */

/**
 * Source: Udemy Go (Golang) The Complete Bootcamp Course
 */

func main() {
	var (
		sum int
		i   = 1
	)

	for {
		if i > 5 {
			break
		}

		if i%2 != 0 {
			i++
			continue
		}

		sum += i
		i++
	}

	fmt.Println(sum)
}
