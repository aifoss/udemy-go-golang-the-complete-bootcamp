package main

import (
	"fmt"
	"os"
)

/**
 * Created by sofia on 2019-12-26.
 */

/**
 * Source: Udemy Go (Golang) The Complete Bootcamp Course
 */

func main() {
	argCount := len(os.Args) - 1

	if argCount == 0 {
		fmt.Println("Give me args")
		return
	}

	if argCount == 1 {
		fmt.Printf("There is one: %q\n", os.Args[1])
	} else if argCount == 2 {
		fmt.Printf("There are two: %q\n", os.Args[1]+" "+os.Args[2])
	} else if argCount > 2 {
		fmt.Printf("There are %d arguments\n", argCount)
	}
}
