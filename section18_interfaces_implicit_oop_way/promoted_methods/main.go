package main

/**
 * Created by sofia on 2019-12-29.
 */

/**
 * Source: Udemy Go (Golang) The Complete Bootcamp Course
 */

func main() {
	store := list{
		&book{product{"moby dick", 10}, 118281600},
		&book{product{"odyssey", 15}, "733622400"},
		&book{product{"hobbit", 25}, nil},
		&puzzle{product{"rubik cube", 5}},
		&game{product{"minecraft", 20}},
		&game{product{"tetris", 10}},
		&toy{product{"yoda", 30}},
	}

	store.discount(.5)

	store.print()
}
