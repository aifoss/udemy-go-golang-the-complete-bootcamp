package main

import "fmt"

/**
 * Created by sofia on 2019-12-23.
 */

/**
 * Source: Udemy Go (Golang) The Complete Bootcamp Course
 */

func main() {
	var apple int
	var orange int32

	apple = int(orange)
	fmt.Println(apple)

	orange = int32(apple)
	fmt.Println(orange)

	orange = 65
	color := string(orange)
	fmt.Println(color)
}
