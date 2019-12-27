package main

import "fmt"

/**
 * Created by sofia on 2019-12-27.
 */

/**
 * Source: Udemy Go (Golang) The Complete Bootcamp Course
 */

func main() {
	blue := [5]int{6, 9, 3, 2, 1}
	red := [5]int{6, 9, 3, 2, 1}

	fmt.Printf("Are they equal? %t\n", blue == red)
	fmt.Printf("blue: %#v\n", blue)
	fmt.Printf("red: %#v\n", red)
	fmt.Println()

	type (
		bookcase [5]int
		cabinet  [5]int
	)

	blue = bookcase{6, 9, 3, 2, 1}
	red = cabinet{6, 9, 3, 2, 1}

	fmt.Printf("Are they equal? %t\n", blue == red)
	fmt.Printf("%#v\n", blue)
	fmt.Printf("%#v\n", cabinet(blue))
	fmt.Printf("%#v\n", bookcase(red))
	fmt.Println()

	type three [3]int

	nums := [3]int{1, 2, 3}
	nums2 := three{1, 2, 3}

	fmt.Println(nums == nums2)

	type (
		threeA [3]int
		threeB [3]int
	)

	nums3 := threeA{1, 2, 3}
	nums4 := threeA(threeB{1, 2, 3})

	fmt.Println(nums3 == nums4)
}
