package main

import "fmt"

/**
 * Created by sofia on 2019-12-29.
 */

/**
 * Source: Udemy Go (Golang) The Complete Bootcamp Course
 */

type book struct {
	title string
	price money
}

func (b book) print() {
	fmt.Printf("%-15s: %s\n", b.title, b.price.string())
}
