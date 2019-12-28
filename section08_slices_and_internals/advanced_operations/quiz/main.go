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
	s.PrintBacking = true

	lyric := []string{"show", "me", "my", "silver", "lining"}
	part := lyric[:2:2]
	part = append(part, "right", "place")

	s.Show("lyric", lyric)
	s.Show("part", part)

	tasks := make([]string, 2)
	tasks = append(tasks, "hello", "world")

	s.Show("tasks", tasks)
	fmt.Printf("%q\n", tasks)

	lyric = []string{"le", "vent", "nous", "portera"}
	n := copy(lyric, make([]string, 4))

	s.Show("lyric", lyric)
	fmt.Printf("%d %q\n", n, lyric)

	spendings := [][]int{{200, 100}, {}, {50, 25, 75}, {500}}
	total := spendings[2][1] + spendings[3][0] + spendings[0][0]

	fmt.Printf("%d\n", total)

	spendings = [][]int{{1, 2}}

	fmt.Printf("%T - ", spendings)
	fmt.Printf("%T - ", spendings[0])
	fmt.Printf("%T\n", spendings[0][0])
}
