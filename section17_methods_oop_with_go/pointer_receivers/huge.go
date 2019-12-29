package main

import "fmt"

/**
 * Created by sofia on 2019-12-29.
 */

/**
 * Source: Udemy Go (Golang) The Complete Bootcamp Course
 */

type huge struct {
	games [10000000]game
}

func (h *huge) addr() {
	fmt.Printf("%p\n", h)
}
