package main

import (
	"fmt"
	"time"
)

/**
 * Created by sofia on 2019-12-26.
 */

/**
 * Source: Udemy Go (Golang) The Complete Bootcamp Course
 */

func main() {
	switch hour := time.Now().Hour(); {
	case hour >= 18:
		fmt.Println("Good evening!")
	case hour >= 12:
		fmt.Println("Good afternoon!")
	case hour >= 6:
		fmt.Println("Good morning!")
	default:
		fmt.Println("Good night!")
	}
}
