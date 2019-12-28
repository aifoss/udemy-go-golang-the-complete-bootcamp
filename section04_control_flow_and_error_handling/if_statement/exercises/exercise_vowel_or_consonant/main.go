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

const vowels = "aeiou"

func main() {
	argCount := len(os.Args) - 1

	if argCount == 0 {
		fmt.Println("Give me a letter")
		return
	}

	input := os.Args[1]
	len := len(input)

	if len > 1 {
		fmt.Println("Give me a letter")
		return
	}

	if strings.IndexAny(input, vowels) >= 0 {
		fmt.Printf("%q is a vowel\n", input)
	} else if input == "y" || input == "w" {
		fmt.Printf("%q is sometimes a vowel, sometimes not\n", input)
	} else {
		fmt.Printf("%q is a consonant\n", input)
	}
}
