package main

import (
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
// Steps:
//
// Check whether there's a command line argument or not.
// If not, quit from the program with a message.
//
// Create a byte buffer as big as the argument.
//
// Loop and detect the http:// patterns
//
// Copy the input character by character to the buffer
//
// If you detect http:// pattern, copy the http:// first,
// then copy the *s instead of the original link until you see whitespace character.
//
// For example:
//
// INPUT:
// Here: http://www.mylink.com Click!
//
// OUTPUT:
// Here: http://************** Click!
// Print the buffer as a string
// ---------------------------------------------------------

func main() {
	const (
		link     = "http://"
		linkSize = len(link)

		mask = '*'
	)

	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("Give me something to mask!")
		return
	}

	var (
		text = args[0]
		size = len(text)
		// buf  = make([]byte, 0, size)
		buf = []byte(text)

		insideLink = false
	)

	for i := 0; i < size; i++ {
		if len(text[i:]) >= linkSize && text[i:i+linkSize] == link {
			insideLink = true
			// buf = append(buf, link...)
			i += linkSize
		}

		c := text[i]

		switch c {
		case ' ', '\t', '\n':
			insideLink = false
		}

		if insideLink {
			// c = mask
			buf[i] = mask
		}

		// buf = append(buf, c)

	}

	fmt.Println(string(buf))
}
