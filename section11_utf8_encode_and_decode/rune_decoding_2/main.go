package main

import (
	"bytes"
	"fmt"
	"unicode/utf8"
)

/**
 * Created by sofia on 2019-12-28.
 */

/**
 * Source: Udemy Go (Golang) The Complete Bootcamp Course
 */

func main() {
	word := []byte("öykü")
	fmt.Printf("%s = % [1]x\n", word)

	// how to make the first rune uppercase?

	var size int

	// for i := range string(word) {
	// 	if i > 0 {
	// 		size = i
	// 		break
	// 	}
	// }

	_, size = utf8.DecodeRune(word)

	copy(word[:size], bytes.ToUpper(word[:size]))

	fmt.Printf("%s = % [1]x\n", word)
}
