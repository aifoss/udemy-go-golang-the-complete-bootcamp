package main

import (
	"fmt"
	"strconv"
)

/**
 * Created by sofia on 2019-12-24.
 */

/**
 * Source: Udemy Go (Golang) The Complete Bootcamp Course
 */

func main() {
	name, last := "carl", "sagan"
	name += " edward"
	fmt.Println(name + " " + last)

	fmt.Println(
		"hello" + ", " + "how" + " " + "are you" + " " +
			"today?",
	)

	fmt.Println(
		`hello` + `, ` + `how` + ` ` + `are you` + ` ` +
			"today?",
	)

	eq := "1 + 2 = "
	sum := 1 + 2

	fmt.Println(eq + strconv.Itoa(sum))

	q := strconv.FormatBool(true) + " or " + strconv.FormatBool(false) + "?"

	fmt.Println(q)
}
