package main

import (
	"fmt"
	"os"
)

/**
 * Created by sofia on 2019-12-28.
 */

/**
 * Source: Udemy Go (Golang) The Complete Bootcamp Course
 */

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("[english word] -> [turkish word]")
		return
	}

	query := args[0]

	// english := []string{"good", "great", "perfect"}
	// turkish := []string{"iyi", "harika", "mükemmel"}

	// for i, w := range english {
	// 	if query == w {
	// 		fmt.Printf("%q means %q\n", w, turkish[i])
	// 		return
	// 	}
	// }

	// fmt.Printf("%q not found\n", query)

	// var dict map[string]string
	dict := map[string]string{
		"good":    "kötü",
		"great":   "harika",
		"perfect": "mükemmel",
	}

	dict["up"] = "yukarı"
	dict["down"] = "aşağı"
	dict["good"] = "iyi"
	dict["mistake"] = ""

	fmt.Printf("%#v\n", dict)

	copied := map[string]string{"down": "aşağı", "good": "iyi", "great": "harika", "mistake": "", "perfect": "mükemmel", "up": "yukarı"}

	first := fmt.Sprintf("%s", dict)
	second := fmt.Sprintf("%s", copied)

	if first == second {
		fmt.Println("Maps are equal")
	}

	if value, ok := dict[query]; ok {
		fmt.Printf("%q means %q\n", query, value)
		return
	}

	fmt.Printf("%q not found\n", query)
}
