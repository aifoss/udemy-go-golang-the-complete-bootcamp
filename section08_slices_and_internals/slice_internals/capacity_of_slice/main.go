package main

import s "github.com/inancgumus/prettyslice"

/**
 * Created by sofia on 2019-12-27.
 */

/**
 * Source: Udemy Go (Golang) The Complete Bootcamp Course
 */

func main() {
	s.PrintBacking = true

	var games []string
	s.Show("games", games)

	games = []string{}
	s.Show("games", games)

	games2 := []string{}
	s.Show("games2", games2)

	games3 := []string{"pacman", "mario", "tetris"}
	s.Show("games3", games3)

	games4 := games3
	s.Show("games4", games4)

	games5 := games4[:0]
	s.Show("games5", games5)
	s.Show("games5", games5[:cap(games5)])

	for cap(games5) != 0 {
		games5 = games5[1:cap(games5)]
		s.Show("games5", games5)
	}
}
