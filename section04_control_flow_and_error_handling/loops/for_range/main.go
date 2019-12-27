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
	for i, v := range os.Args {
		if i == 0 {
			continue
		}
		fmt.Printf("%q\n", v)
	}

	for _, v := range os.Args[1:] {
		fmt.Printf("%q\n", v)
	}

	words := strings.Fields(
		"lazy cat jumps again and again and again",
	)

	var (
		i int
		v string
	)

	for i, v = range words {
		fmt.Printf("#%-2d: %q\n", i, v)
	}

	fmt.Printf("Last value: i = %d, v = %q\n", i, v)
}
