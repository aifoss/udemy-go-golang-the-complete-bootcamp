package main

import "fmt"

/**
 * Created by sofia on 2019-12-29.
 */

/**
 * Source: Udemy Go (Golang) The Complete Bootcamp Course
 */

type printer interface {
	print()
}

type list []printer

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
	type discounter interface {
		discount(float64)
	}

	for _, it := range l {
		// g, isGame := it.(*game)
		// if !isGame {
		// 	continue
		// }
		// g.discount(ratio)

		// i, ok := it.(interface{ discount(float64) })
		// if !ok {
		// 	continue
		// }
		// i.discount(ratio)

		if it, ok := it.(discounter); ok {
			it.discount(ratio)
		}
	}
}
