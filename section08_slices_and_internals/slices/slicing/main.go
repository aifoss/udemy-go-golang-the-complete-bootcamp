package main

import (
	"fmt"

	s "github.com/inancgumus/prettyslice"
)

/**
 * Created by sofia on 2019-12-27.
 */

/**
 * Source: Udemy Go (Golang) The Complete Bootcamp Course
 */

func main() {
	items := []string{
		"pacman", "mario", "tetris", "doom",
		"galaga", "frogger", "asteroids", "simcity",
		"metroid", "defender", "rayman", "tempest",
		"ultima",
	}

	s.MaxPerLine = 4
	s.Show("items", items)

	top3 := items[:3]

	s.Show("top3", top3)

	l := len(items)
	last4 := items[l-4:]

	s.Show("last4", last4)

	mid := last4[1:3]

	s.Show("mid", mid)

	fmt.Printf("slicing: %T[ %[1]q\n", items[2:3])
	fmt.Printf("indexing: %T[ %[1]q\n", items[2])
}
