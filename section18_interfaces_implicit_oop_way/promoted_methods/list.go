package main

import "fmt"

/**
 * Created by sofia on 2019-12-29.
 */

/**
 * Source: Udemy Go (Golang) The Complete Bootcamp Course
 */

type item interface {
	print()
	discount(ratio float64)
}

type list []item

func (l list) print() {
	if len(l) == 0 {
		fmt.Println("Sorry. We're waiting for delivery.")
		return
	}

	for _, i := range l {
		i.print()
	}
}

func (l list) discount(ratio float64) {
	for _, it := range l {
		it.discount(ratio)
	}
}
