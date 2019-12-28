package main

import (
	"fmt"
	"strconv"
	"strings"
)

/**
 * Created by sofia on 2019-12-27.
 */

/**
 * Source: Udemy Go (Golang) The Complete Bootcamp Course
 */

func main() {
	spendings := [][]int{
		{200, 100},
		{500},
		{50, 25, 75},
	}

	printSpendings(spendings)

	spendings = make([][]int, 0, 5)

	spendings = append(spendings, []int{200, 100})
	spendings = append(spendings, []int{500})
	spendings = append(spendings, []int{50, 25, 75})

	printSpendings(spendings)

	spendings = fetch()

	printSpendings(spendings)
}

func fetch() [][]int {
	content := `200 100
25 10 45 60
5 15 35
95 10
50 25`

	lines := strings.Split(content, "\n")

	spendings := make([][]int, len(lines))

	for i, line := range lines {
		fmt.Printf("%d: %#v\n", i+1, line)

		fields := strings.Fields(line)

		spendings[i] = make([]int, len(fields))

		for j, field := range fields {
			fmt.Printf("\t%d: %#v\n", j+1, field)

			spending, _ := strconv.Atoi(field)

			spendings[i][j] = spending
		}
	}

	fmt.Println()

	return spendings
}

func printSpendings(spendings [][]int) {
	for i, daily := range spendings {
		var total int

		for _, spending := range daily {
			total += spending
		}

		fmt.Printf("Day %d: %d\n", i+1, total)
	}

	fmt.Println()
}
