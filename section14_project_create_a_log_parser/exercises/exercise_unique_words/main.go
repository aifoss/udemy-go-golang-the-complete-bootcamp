package main

import (
	"bufio"
	"fmt"
	"os"
)

/**
 * Created by sofia on 2019-12-28.
 */

/**
 * Source: Udemy Go (Golang) The Complete Bootcamp Course
 */

// ---------------------------------------------------------
// EXERCISE: Unique Words
//
//  Create a program that prints the total and unique words
//  from an input stream.
//
//  1. Feed the shakespeare.txt to your program.
//
//  2. Scan the input using a new Scanner.
//
//  3. Configure the scanner to scan for the words.
//
//  4. Count the unique words using a map.
//
//  5. Print the total and unique words.
//
//
// EXPECTED OUTPUT
//
//  There are 99 words, 70 of them are unique.
//
// ---------------------------------------------------------

func main() {
	counts := 0
	uniques := make(map[string]int)

	in := bufio.NewScanner(os.Stdin)
	in.Split(bufio.ScanWords)

	for in.Scan() {
		word := in.Text()
		uniques[word]++
		counts++
	}

	fmt.Printf("Total Words: %d\nUnique Words: %d\n", counts, len(uniques))
}
