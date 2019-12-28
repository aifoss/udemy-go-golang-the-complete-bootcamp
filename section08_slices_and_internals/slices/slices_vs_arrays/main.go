package main

import "fmt"

/**
 * Created by sofia on 2019-12-27.
 */

/**
 * Source: Udemy Go (Golang) The Complete Bootcamp Course
 */

func main() {
	var books [5]string
	books[0] = "dracula"
	books[1] = "1984"
	books[2] = "island"

	fmt.Printf("books: %#v\n", books)
	fmt.Println()

	books2 := [5]string{"ulysses", "fire"}

	fmt.Printf("books2: %#v\n", books2)
	fmt.Println()

	var games []string

	fmt.Printf("games: %#v\n", games)
	fmt.Printf("len(games): %d\n", len(games))
	fmt.Printf("nil? :%t\n", games == nil)
	fmt.Println()

	games2 := []string{"kokemon", "sims"}

	fmt.Printf("games2: %#v\n", games2)
	fmt.Printf("len(games2): %d\n", len(games2))
	fmt.Printf("nil? :%t\n", games2 == nil)
	fmt.Println()

	games3 := []string{"pacman", "doom", "pong"}

	var ok string

	if len(games) != len(games3) {
		ok = "not"
	}

	fmt.Printf("games: %#v\n", games)
	fmt.Printf("games3: %#v\n", games3)
	fmt.Printf("games and games3 are %s equal\n", ok)
	fmt.Println()

	ok = ""

	for i, game := range games2 {
		if game != games3[i] {
			ok = "not"
			break
		}
	}

	fmt.Printf("games2: %#v\n", games2)
	fmt.Printf("games3: %#v\n", games3)
	fmt.Printf("games2 and games3 are %s equal\n", ok)
	fmt.Println()

	games3 = games2

	ok = ""

	for i, game := range games2 {
		if game != games3[i] {
			ok = "not"
			break
		}
	}

	fmt.Printf("games2: %#v\n", games2)
	fmt.Printf("games3: %#v\n", games3)
	fmt.Printf("games2 and games3 are %s equal\n", ok)
}
