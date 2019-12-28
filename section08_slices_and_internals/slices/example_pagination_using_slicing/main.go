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

	const pageSize = 4

	l := len(items)

	for from := 0; from < l; from += pageSize {
		to := from + pageSize

		if to > l {
			to = l
		}

		head := fmt.Sprintf("Page #%d", (from/pageSize)+1)
		currPage := items[from:to]

		s.Show(head, currPage)
	}
}
