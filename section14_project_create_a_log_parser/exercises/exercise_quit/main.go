package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/**
 * Created by sofia on 2019-12-28.
 */

/**
 * Source: Udemy Go (Golang) The Complete Bootcamp Course
 */

// ---------------------------------------------------------
// EXERCISE: Quit
//
//  Create a program that quits when a user types the
//  same word twice.
//
//
// RESTRICTION
//
//  The program should work case insensitive.
//
//
// EXPECTED OUTPUT
//
//  go run main.go
//
//   hey
//   HEY
//   TWICE!
//
//  go run main.go
//
//   hey
//   hi
//   HEY
//   TWICE!
// ---------------------------------------------------------

func main() {
	words := make(map[string]bool)

	in := bufio.NewScanner(os.Stdin)
	in.Split(bufio.ScanWords)

	for in.Scan() {
		word := strings.ToLower(in.Text())

		if words[word] {
			fmt.Println("TWICE!")
			return
		}

		words[word] = true
	}
}
