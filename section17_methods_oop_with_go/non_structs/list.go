package main

import "fmt"

/**
 * Created by sofia on 2019-12-29.
 */

/**
 * Source: Udemy Go (Golang) The Complete Bootcamp Course
 */

type list []*game

func (l list) print() {
	if len(l) == 0 {
		fmt.Println("Sorry. We're waiting for delivery.")
		return
	}

	for _, i := range l {
		i.print()
	}
}
