package main

import (
	"strings"

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
	s.MaxPerLine = 10

	tasks := []string{"jump", "run", "read"}

	var upTasks []string
	s.Show("upTasks", upTasks)

	for _, task := range tasks {
		// allocates new backing array each time
		upTasks = append(upTasks, strings.ToUpper(task))
		s.Show("upTasks", upTasks)
	}

	efficientUpTasks := make([]string, 0, len(tasks))
	s.Show("effUpTasks", efficientUpTasks)

	// for i := range tasks {
	// 	efficientUpTasks[i] = strings.ToUpper(tasks[i])
	// 	s.Show("effUpTasks", efficientUpTasks)
	// }

	for _, task := range tasks {
		efficientUpTasks = append(efficientUpTasks, strings.ToUpper(task))
		s.Show("effUpTasks", efficientUpTasks)
	}
}
