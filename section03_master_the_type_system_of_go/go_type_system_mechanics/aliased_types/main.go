package main

/**
 * Created by sofia on 2019-12-25.
 */

/**
 * Source: Udemy Go (Golang) The Complete Bootcamp Course
 */

// type byte = uint8
// type rune = int32

func main() {
	var (
		byteVal  byte
		uint8Val uint8
	)

	uint8Val = byteVal
	byteVal = uint8Val

	var (
		runeVal  rune
		int32Val int32
	)

	runeVal = int32Val
	int32Val = runeVal
}
