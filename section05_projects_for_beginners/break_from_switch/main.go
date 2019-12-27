package main

import (
	"fmt"
	"os"
	"strings"
)

/**
 * Created by sofia on 2019-12-26.
 */

/**
 * Source: Udemy Go (Golang) The Complete Bootcamp Course
 */

const corpus = "" +
	"lazy cat jumps again and again and again"

func main() {
	words := strings.Fields(corpus)

	query := os.Args[1:]

queries:
	for _, q := range query {

	search:
		for i, w := range words {
			switch q {
			case "and", "or", "the":
				break search
			}

			if q == w {
				fmt.Printf("#%-2d: %q\n",
					i+1, w)
				break queries
			}
		}
	}
}
