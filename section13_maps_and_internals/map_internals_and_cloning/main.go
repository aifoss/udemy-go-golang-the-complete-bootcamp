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

	dict := map[string]string{
		"good":    "kötü",
		"great":   "harika",
		"perfect": "mükemmel",
		"awesome": "mükemmel",
	}

	dict["up"] = "yukarı"
	dict["down"] = "aşağı"
	dict["good"] = "iyi"

	delete(dict, "awesome")

	turkish := make(map[string]string, len(dict))

	for k, v := range dict {
		turkish[v] = k
	}

	fmt.Printf("english: %q\nturkish: %q\n", dict, turkish)

	if value, ok := turkish[query]; ok {
		fmt.Printf("%q means %q\n", query, value)
		return
	}

	if value, ok := dict[query]; ok {
		fmt.Printf("%q means %q\n", query, value)
		return
	}

	fmt.Printf("%q not found\n", query)
}
