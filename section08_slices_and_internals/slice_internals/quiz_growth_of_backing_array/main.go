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
	s.MaxPerLine = 30
	s.Width = 150

	words := []string{"lucy", "in", "the", "sky", "with", "diamonds"}

	s.Show("words", words)

	words = append(words[:1], "is", "everywhere")

	s.Show("words", words)

	words = append(words, words[len(words)+1:cap(words)]...)

	s.Show("words", words)
}
