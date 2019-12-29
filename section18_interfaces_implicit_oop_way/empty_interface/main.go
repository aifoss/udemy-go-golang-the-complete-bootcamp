package main

/**
 * Created by sofia on 2019-12-29.
 */

/**
 * Source: Udemy Go (Golang) The Complete Bootcamp Course
 */

func main() {
	store := list{
		book{title: "moby dick", price: 10, published: 118281600},
		book{title: "odyssey", price: 15, published: "733622400"},
		book{title: "hobbit", price: 25},
		puzzle{title: "rubik cube", price: 5},
		&game{title: "minecraft", price: 20},
		&game{title: "tetris", price: 10},
		&toy{title: "yoda", price: 30},
	}

	store.discount(.5)

	store.print()
}
