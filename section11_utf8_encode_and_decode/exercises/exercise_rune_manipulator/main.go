package main

import (
	"fmt"
	"unicode/utf8"
)

/**
 * Created by sofia on 2019-12-28.
 */

/**
 * Source: Udemy Go (Golang) The Complete Bootcamp Course
 */

// ---------------------------------------------------------
// EXERCISE: Rune Manipulator
//
//  Please read the comments inside the following code.
//
// EXPECTED OUTPUT
//  Please run the solution.
// ---------------------------------------------------------

func main() {
	strings := []string{
		"cool",
		"güzel",
		"jīntiān",
		"今天",
		"read 🤓",
	}

	for _, s := range strings {
		fmt.Printf("%q\n", s)

		// Print the byte and rune length of the strings
		// Hint: Use len and utf8.RuneCountInString
		fmt.Printf("\thas %d bytes and %d runes\n",
			len(s), utf8.RuneCountInString(s))

		// Print the bytes of the strings in hexadecimal
		// Hint: Use % x verb
		fmt.Printf("\tbytes: % x\n", s)

		// Print the runes of the strings in hexadecimal
		// Hint: Use % x verb
		fmt.Print("\trune hexes: ")
		for _, r := range s {
			fmt.Printf("% x", r)
		}
		fmt.Println()

		// Print the runes of the strings as rune literals
		// Hint: Use for range
		fmt.Print("\trune literals: ")
		for _, r := range s {
			fmt.Printf("%q", r)
		}
		fmt.Println()

		// Print the first rune and its byte size of the strings
		// Hint: Use utf8.DecodeRuneInString
		r, size := utf8.DecodeRuneInString(s)
		fmt.Printf("\tfirst: %q (%d bytes)\n", r, size)

		// Print the last rune of the strings
		// Hint: Use utf8.DecodeLastRuneInString
		r, size = utf8.DecodeLastRuneInString(s)
		fmt.Printf("\tlast: %q (%d bytes)\n", r, size)

		// Slice and print the first two runes of the strings
		_, first1 := utf8.DecodeRuneInString(s)
		_, first2 := utf8.DecodeRuneInString(s[first1:])
		fmt.Printf("\tfirst 2: %q\n", s[:first1+first2])

		// Slice and print the last two runes of the strings
		_, last1 := utf8.DecodeLastRuneInString(s)
		_, last2 := utf8.DecodeLastRuneInString(s[:len(s)-1])
		fmt.Printf("\tlast 2: %q\n", s[len(s)-last2-last1:])

		// Convert the string to []rune
		// Print the first and last two runes
		rs := []rune(s)
		fmt.Printf("\tfirst 2: %q\n", string(rs[:2]))
		fmt.Printf("\tlast 2: %q\n", string(rs[len(rs)-2:]))
	}
}
