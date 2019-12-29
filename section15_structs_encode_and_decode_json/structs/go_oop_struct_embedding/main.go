package main

import "fmt"

/**
 * Created by sofia on 2019-12-28.
 */

/**
 * Source: Udemy Go (Golang) The Complete Bootcamp Course
 */

func main() {
	type text struct {
		title string
		words int
	}

	type book struct {
		text // anonymous field
		isbn string
	}

	moby := book{
		text: text{title: "moby dick", words: 206052},
		isbn: "102030",
	}

	fmt.Printf("%s has %d words (isbn: %s)\n",
		// moby.text.title, moby.text.words, moby.isbn,
		moby.title, moby.words, moby.isbn,
	)
}
