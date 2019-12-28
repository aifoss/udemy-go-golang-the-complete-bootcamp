package main

import "fmt"

/**
 * Created by sofia on 2019-12-26.
 */

/**
 * Source: Udemy Go (Golang) The Complete Bootcamp Course
 */

func main() {
	const home = "Milky Way Galaxy"
	const length = len(home)

	fmt.Printf("There are %d characters inside %q\n", length, home)
}
