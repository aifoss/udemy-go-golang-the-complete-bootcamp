package main

import "fmt"

/**
 * Created by sofia on 2019-12-27.
 */

/**
 * Source: Udemy Go (Golang) The Complete Bootcamp Course
 */

// ---------------------------------------------------------
// EXERCISE: Observe the length and capacity
//
//  Follow the instructions inside the code below to
//  gain more intuition about the length and capacity of a slice.
//
// ---------------------------------------------------------

func main() {
	// --- #1 ---
	// 1. create a new slice named: games
	// var games []string

	// 2. print the length and capacity of the games slice
	// fmt.Printf("games's len    : %d cap: %d\n", len(games), cap(games))

	// 3. comment out the games slice
	//    then declare it as an empty slice
	// games := []string{}

	// 4. print the length and capacity of the games slice
	// fmt.Printf("games's len    : %d cap: %d\n", len(games), cap(games))

	// 5. append the elements: "pacman", "mario", "tetris", "doom"
	// games = append(games, "pacman", "mario", "tetris", "doom")

	// 6. print the length and capacity of the games slice
	// fmt.Printf("games's len    : %d cap: %d\n", len(games), cap(games))

	// 7. comment out everything

	// 8. declare the games slice again using a slice literal
	//    (use the same elements from step 3)
	games := []string{"pacman", "mario", "tetris", "doom"}
	fmt.Printf("games's len    : %d cap: %d\n", len(games), cap(games))
	fmt.Println()

	// --- #2 ---
	// 1. use a loop from 0 to 4 to slice the games slice, element by element.
	// 2. print its length and capacity along the way (in the loop).
	for i := 0; i < 4; i++ {
		s := games[:i]
		fmt.Printf("games[:%d]'s len: %d cap: %d\n", i, len(s), cap(s))
	}
	fmt.Println()

	// --- #3 ---
	// 1. slice the games slice up to zero element
	//    (save the result to a new slice named: "zero")
	zero := games[:0]

	// 2. print the games and the new slice's len and cap
	fmt.Printf("games's     len: %d cap: %d\n", len(games), cap(games))
	fmt.Printf("zero's      len: %d cap: %d\n", len(zero), cap(zero))

	// 3. append a new element to the new slice
	// 4. print the new slice's lens and caps
	// 5. repeat the last two steps 5 times (use a loop)
	// 6. notice the growth of the capacity after the 5th append
	for _, v := range []string{"ultima", "dagger", "pong", "coldspot", "zetra"} {
		zero = append(zero, v)
		fmt.Printf("zero's      len: %d cap: %d\n", len(zero), cap(zero))
	}
	fmt.Println()

	// --- #4 ---
	// using a range loop, slice the zero slice element by element,
	// and print its length and capacity along the way.
	//
	// observe that, the range loop only loops for the length, not the cap.
	for n := range zero {
		s := zero[:n]
		fmt.Printf("zero[:%d]'s  len: %d cap: %d\n", n, len(s), cap(s))
	}
	fmt.Println()

	// --- #5 ---
	// 1. do the 3rd step above again but this time, start by slicing
	//    the zero slice up to its capacity (use the cap function).
	//
	// 2. print the elements of the zero slice in the loop.
	zero = zero[:cap(zero)]
	for n := range zero {
		s := zero[:n]
		fmt.Printf("zero[:%d]'s  len: %d cap: %d - %q\n", n, len(s), cap(s), s)
	}
	fmt.Println()

	// --- #6 ---
	// 1. change the one of the elements of the zero slice
	zero[0] = "command & conquer"

	// 2. change the same element of the games slice
	games[0] = "red alert"

	// 3. print the games and the zero slices
	fmt.Printf("zero  : %q\n", zero)
	fmt.Printf("games : %q\n", games)
	fmt.Println()

	// 4. observe that they don't have the same backing array

	// --- #7 ---
	// try to slice the games slice beyond its capacity
	// uncomment and see the error.
	// _ = games[:cap(games)+1]
	// or:
	// _ = games[:5]
}
