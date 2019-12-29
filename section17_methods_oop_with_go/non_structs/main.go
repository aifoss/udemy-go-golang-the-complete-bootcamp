package main

/**
 * Created by sofia on 2019-12-29.
 */

/**
 * Source: Udemy Go (Golang) The Complete Bootcamp Course
 */

func main() {
	var (
		// mobydick  = book{title: "moby dick", price: 10}
		minecraft = game{title: "minecraft", price: 20}
		tetris    = game{title: "tetris", price: 10}
	)

	var items []*game

	items = append(items, &minecraft, &tetris)

	ml := list(items)

	ml.print()
}
