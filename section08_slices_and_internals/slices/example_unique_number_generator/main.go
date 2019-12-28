package main

import (
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"time"
)

/**
 * Created by sofia on 2019-12-27.
 */

/**
 * Source: Udemy Go (Golang) The Complete Bootcamp Course
 */

func main() {
	rand.Seed(time.Now().UnixNano())

	max, _ := strconv.Atoi(os.Args[1])

	// var uniques [max]int
	var uniques []int

	// loop:
	// 	for found := 0; found < max; {
	// 		n := rand.Intn(max) + 1
	// 		fmt.Print(n, " ")

	// 		for _, u := range uniques {
	// 			if u == n {
	// 				continue loop
	// 			}
	// 		}

	// 		uniques[found] = n
	// 		found++
	// 	}

loop:
	for len(uniques) < max {
		n := rand.Intn(max) + 1
		fmt.Print(n, " ")

		for _, u := range uniques {
			if u == n {
				continue loop
			}
		}

		uniques = append(uniques, n)
	}

	fmt.Println("\n\nunqiues:", uniques)

	sort.Ints(uniques)
	fmt.Println("\nsorted:", uniques)

	nums := [5]int{5, 4, 3, 2, 1}
	sort.Ints(nums[:])
	fmt.Println("\nnums:", nums)
}
