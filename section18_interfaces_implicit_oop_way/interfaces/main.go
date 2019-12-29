package main

/**
 * Created by sofia on 2019-12-29.
 */

/**
 * Source: Udemy Go (Golang) The Complete Bootcamp Course
 */

func main() {
	var (
		mobydick  = book{title: "moby dick", price: 10}
		minecraft = game{title: "minecraft", price: 20}
		tetris    = game{title: "tetris", price: 10}
		rubik     = puzzle{title: "rubik cube", price: 5}
		yoda      = toy{title: "yoda", price: 30}
	)

	var store list
	store = append(store, &minecraft, &tetris, mobydick, rubik, &yoda)

	// fmt.Println(store[0] == &minecraft)
	// fmt.Println(store[3] == rubik)

	// var p printer
	// p = mobydick
	// p.print()
	// p = &tetris
	// p.print()

	store.discount(.5)

	store.print()
}
