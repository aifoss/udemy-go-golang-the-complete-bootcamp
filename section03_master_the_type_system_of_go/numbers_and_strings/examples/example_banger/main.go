package main

import (
	"fmt"
	"os"
	"strings"
)

/**
 * Created by sofia on 2019-12-24.
 */

/**
 * Source: Udemy Go (Golang) The Complete Bootcamp Course
 */

func main() {
	msg := os.Args[1]
	len := len(msg)
	bang := strings.Repeat("!", len)
	s := bang + msg + bang
	s = strings.ToUpper(s)

	fmt.Println(s)
}
