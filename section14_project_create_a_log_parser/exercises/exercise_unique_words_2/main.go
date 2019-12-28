package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

/**
 * Created by sofia on 2019-12-28.
 */

/**
 * Source: Udemy Go (Golang) The Complete Bootcamp Course
 */

// ---------------------------------------------------------
// EXERCISE: Unique Words 2
//
//  Use your solution from the previous "Unique Words"
//  exercise.
//
//  Before adding the words to your map, remove the
//  punctuation characters and numbers from them.
//
//
// BE CAREFUL
//
//  Now the shakespeare.txt contains upper and lower
//  case letters too.
//
//
// EXPECTED OUTPUT
//
//  go run main.go < shakespeare.txt
//
//   There are 100 words, 69 of them are unique.
//
// ---------------------------------------------------------

func main() {
	rx := regexp.MustCompile(`[^A-Za-z]+`)

	counts := 0
	uniques := make(map[string]int)

	in := bufio.NewScanner(os.Stdin)
	in.Split(bufio.ScanWords)

	for in.Scan() {
		word := in.Text()
		word = rx.ReplaceAllString(word, "")

		if word != "" {
			word = strings.ToLower(word)
			uniques[word]++
			counts++
		}
	}

	fmt.Printf("There are %d words, %d of them are unique.\n", counts, len(uniques))
}
