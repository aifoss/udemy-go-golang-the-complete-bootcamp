package main

import (
	"fmt"
	"strconv"
)

/**
 * Created by sofia on 2019-12-26.
 */

/**
 * Source: Udemy Go (Golang) The Complete Bootcamp Course
 */

func main() {
	if n, err := strconv.Atoi("42"); err == nil {
		fmt.Println("There was no error, n is", n)
	}
}
