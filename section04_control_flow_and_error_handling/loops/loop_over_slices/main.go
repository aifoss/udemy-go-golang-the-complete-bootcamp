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

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("%q\n", os.Args[i])
	}

	words := strings.Fields(
		"lazy cat jumps again and again and again",
	)

	for i := 0; i < len(words); i++ {
		fmt.Printf("#%-2d: %q\n", i, words[i])
	}
}
