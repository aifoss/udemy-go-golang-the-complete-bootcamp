package main

import (
	"bytes"
	"fmt"
)

/**
 * Created by sofia on 2019-12-27.
 */

/**
 * Source: Udemy Go (Golang) The Complete Bootcamp Course
 */

// ---------------------------------------------------------
// EXERCISE: Append
//
//  Please follow the instructions within the code below.
//
// EXPECTED OUTPUT
//  They are equal.
//
// HINTS
//  bytes.Equal function allows you to compare two byte
//  slices easily. Check its documentation: go doc bytes.Equal
// ---------------------------------------------------------

func main() {
	// 1. uncomment the code below
	png, header := []byte{'P', 'N', 'G'}, []byte{}

	// 2. append elements to header to make it equal with the png slice
	header = append(header, png...)

	// 3. compare the slices using the bytes.Equal function
	equal := bytes.Equal(png, header)

	// 4. print whether they're equal or not
	switch equal {
	case true:
		fmt.Println("They are equal.")
	case false:
		fmt.Println("They are NOT equal.")
	}
}
