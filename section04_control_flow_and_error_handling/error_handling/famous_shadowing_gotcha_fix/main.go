package main

import (
	"fmt"
	"os"
	"strconv"
)

/**
 * Created by sofia on 2019-12-26.
 */

/**
 * Source: Udemy Go (Golang) The Complete Bootcamp Course
 */

func main() {
	var (
		n   int
		err error
	)

	if a := os.Args; len(a) != 2 {
		fmt.Println("Give me a number")
	} else if n, err = strconv.Atoi(a[1]); err != nil {
		fmt.Printf("Cannot convert %q\n", a[1])
	} else {
		n *= 2
		fmt.Printf("%s * 2 = %d\n", a[1], n)
	}

	fmt.Printf("n is %d\n", n)
}
