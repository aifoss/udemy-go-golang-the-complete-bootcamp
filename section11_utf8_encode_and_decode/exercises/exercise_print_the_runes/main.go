package main

import (
	"fmt"
)

/**
 * Created by sofia on 2019-12-28.
 */

/**
 * Source: Udemy Go (Golang) The Complete Bootcamp Course
 */

// ---------------------------------------------------------
// EXERCISE: Print the runes
//
//  1. Loop over the "console" word and print its runes one by one,
//     in decimals, hexadecimals and binary.
//
//  2. Manually put the runes of the "console" word to a byte slice, one by one.
//
//     As the elements of the byte slice use only the rune literals.
//
//     Print the byte slice.
//
//  3. Repeat the step 2 but this time, as the elements of the byte slice,
//     use only decimal numbers.
//
//  4. Repeat the step 2 but this time, as the elements of the byte slice,
//     use only hexadecimal numbers.
//
//
// EXPECTED OUTPUT
//   Run the solution to see the expected output.
// ---------------------------------------------------------

func main() {
	const word = "console"

	for _, w := range word {
		fmt.Printf("%c\n", w)
		fmt.Printf("\tdecimal: %[1]d\n", w)
		fmt.Printf("\thex: 0x%[1]x\n", w)
		fmt.Printf("\tbinary: 0b%[1]b\n", w)
	}

	fmt.Println()
	fmt.Printf("with runes: %s\n",
		string([]byte{'c', 'o', 'n', 's', 'o', 'l', 'e'}),
	)

	fmt.Println()
	fmt.Printf("with decimals: %s\n",
		string([]byte{99, 111, 110, 115, 111, 100, 101}),
	)

	fmt.Println()
	fmt.Printf("with hexadecimals: %s\n",
		string([]byte{0x63, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65}),
	)
}
