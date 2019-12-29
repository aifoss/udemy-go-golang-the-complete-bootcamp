package main

/**
 * Created by sofia on 2019-12-29.
 */

/**
 * Source: Udemy Go (Golang) The Complete Bootcamp Course
 */

func main() {
	mobydick := book{
		title: "moby dick",
		price: 10,
	}

	mobydick.print()

	minecraft := game{
		title: "minecraft",
		price: 20,
	}

	tetris := game{
		title: "tetris",
		price: 10,
	}

	minecraft.print()
	tetris.print()
}
