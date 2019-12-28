package main

/**
 * Created by sofia on 2019-12-27.
 */

/**
 * Source: Udemy Go (Golang) The Complete Bootcamp Course
 */

import s "github.com/inancgumus/prettyslice"

func main() {
	var todo []string

	s.Show("todo", todo)

	todo = append(todo, "sing")

	s.Show("todo", todo)

	todo = append(todo, "run", "code", "play")

	s.Show("todo", todo)

	tomorrow := []string{"call mom", "learn go"}

	todo = append(todo, tomorrow...)

	s.Show("todo", todo)
}
