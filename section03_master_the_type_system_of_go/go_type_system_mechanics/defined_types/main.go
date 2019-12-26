package main

import "fmt"

/**
 * Created by sofia on 2019-12-25.
 */

/**
 * Source: Udemy Go (Golang) The Complete Bootcamp Course
 */

// Duration is a defined type with int64 as the underlying type
type Duration int64

func main() {
	var ms int64 = 1000
	var ns Duration
	ns = Duration(ms)
	ms = int64(ns)

	fmt.Println(ms)
	fmt.Println(ns)
}
