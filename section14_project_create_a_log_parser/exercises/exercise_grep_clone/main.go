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
// EXERCISE: Grep Clone
//
//  Create a grep clone. grep is a command-line utility for
//  searching plain-text data for lines that match a specific
//  pattern.
//
//  1. Feed the shakespeare.txt to your program.
//
//  2. Accept a command-line argument for the pattern
//
//  3. Only print the lines that contains that pattern
//
//  4. If no pattern is provided, print all the lines
//
//
// EXPECTED OUTPUT
//
//  go run main.go come < shakespeare.txt
//
//    come night come romeo come thou day in night
//    come gentle night come loving black-browed night
//
// ---------------------------------------------------------

func main() {
	args := os.Args[1:]
	pattern := ""

	if len(args) == 1 {
		pattern = args[0]
	}

	in := bufio.NewScanner(os.Stdin)
	in.Split(bufio.ScanLines)

	for in.Scan() {
		line := in.Text()

		if pattern == "" {
			fmt.Println(line)
		} else {
			if strings.Contains(line, pattern) {
				fmt.Println(line)
			}
		}
	}
}
