package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
)

/**
 * Created by sofia on 2019-12-27.
 */

/**
 * Source: Udemy Go (Golang) The Complete Bootcamp Course
 */

// ---------------------------------------------------------
// EXERCISE: Sort and write items to a file
//
//  1. Get arguments from command-line
//
//  2. Sort them
//
//  3. Write the sorted slice to a file
//
//
// EXPECTED OUTPUT
//
//   go run main.go
//     Send me some items and I will sort them
//
//   go run main.go orange banana apple
//
//   cat sorted.txt
//     apple
//     banana
//     orange
//
//
// HINTS
//
//   + REMEMBER: os.Args is a []string
//
//   + String slices are sortable using `sort.Strings`
//
//   + Use ioutil.WriteFile to write to a file.
//
//   + But you need to convert []string to []byte to be able to
//     write it to a file using the ioutil.WriteFile.
//
//   + To do that, create a new []byte and append the elements of your
//     []string.
//
// ---------------------------------------------------------

const maxItems = 10

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Send me some items and I will sort them")
		return
	}

	items := make([]string, 0, maxItems)

	for _, item := range args {
		items = append(items, item)
	}

	sort.Strings(items)

	itemsToWrite := make([]byte, 0, len(items))

	for _, item := range items {
		itemsToWrite = append(itemsToWrite, item...)
		itemsToWrite = append(itemsToWrite, '\n')
	}

	err := ioutil.WriteFile("sorted.txt", itemsToWrite, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Done!")
}
