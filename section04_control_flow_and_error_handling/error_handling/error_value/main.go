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
	s := strconv.Itoa(42)
	fmt.Println(s)

	n, error := strconv.Atoi(os.Args[1])
	fmt.Println("Converted number:", n)
	fmt.Println("Returned error:", error)
}
