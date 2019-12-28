package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

/**
 * Created by sofia on 2019-12-26.
 */

/**
 * Source: Udemy Go (Golang) The Complete Bootcamp Course
 */

// go run main.go 10
func main() {
	t, _ := time.ParseDuration("1h30m")

	multiplier, _ := strconv.ParseInt(os.Args[1], 10, 64)
	t *= time.Duration(multiplier)

	fmt.Println(t)
}
