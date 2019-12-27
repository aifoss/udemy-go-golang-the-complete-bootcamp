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
	age := os.Args[1]
	n, err := strconv.Atoi(age)

	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}

	fmt.Printf("SUCESS: Converted %q to %d\n", age, n)
}
