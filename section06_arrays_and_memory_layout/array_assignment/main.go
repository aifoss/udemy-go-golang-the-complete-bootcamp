package main

import "fmt"

/**
 * Created by sofia on 2019-12-27.
 */

/**
 * Source: Udemy Go (Golang) The Complete Bootcamp Course
 */

func main() {
	prev := [3]string{
		"Kafka's Revenge",
		"Stay Golden",
		"Everythingship",
	}

	curr := prev

	for i := range prev {
		curr[i] += " 2nd Ed."
	}

	fmt.Printf("last year:\n%#v\n", prev)
	fmt.Printf("\nthis year:\n%#v\n", curr)

	var books [4]string

	for i, b := range prev {
		books[i] += b + " 2nd Ed."
	}
	books[3] = "Awesomeness"

	fmt.Printf("\nbooks:\n%#v\n", books)
}
