package main

import (
	"fmt"
	"unicode/utf8"
)

/**
 * Created by sofia on 2019-12-24.
 */

/**
 * Source: Udemy Go (Golang) The Complete Bootcamp Course
 */

func main() {
	name := "carl"

	fmt.Println(len(name))
	fmt.Println(utf8.RuneCountInString(name))
}
