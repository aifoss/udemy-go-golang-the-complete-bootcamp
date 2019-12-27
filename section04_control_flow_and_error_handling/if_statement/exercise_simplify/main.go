package main

import "fmt"

/**
 * Created by sofia on 2019-12-26.
 */

/**
 * Source: Udemy Go (Golang) The Complete Bootcamp Course
 */

func main() {
	isSphere, radius := true, 200

	var big bool

	if radius >= 200 {
		big = true
	}

	if isSphere && big {
		fmt.Println("It's a big sphere")
	} else {
		fmt.Println("I don't know")
	}
}
