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
// EXERCISE: Uppercaser
//
//  Use a scanner to convert the lines to uppercase, and
//  print them.
//
//  1. Feed the shakespeare.txt to your program.
//
//  2. Scan the input using a new Scanner.
//
//  3. Print each line.
//
// EXPECTED OUTPUT
//  Please run the solution to see the expected output.
// ---------------------------------------------------------

func main() {
	in := bufio.NewScanner(os.Stdin)
	in.Split(bufio.ScanLines)

	for in.Scan() {
		line := strings.ToUpper(in.Text())
		fmt.Println(line)
	}
}
