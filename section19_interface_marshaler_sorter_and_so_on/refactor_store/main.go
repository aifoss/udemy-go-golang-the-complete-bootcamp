package main

import (
	"encoding/json"
	"fmt"
	"log"
	"sort"
)

/**
 * Created by sofia on 2019-12-29.
 */

/**
 * Source: Udemy Go (Golang) The Complete Bootcamp Course
 */

const data = `
[
 {
  "title": "hobbit",
  "price": 12.5,
  "released": -62135596800
 },
 {
  "title": "moby dick",
  "price": 5,
  "released": 118281600
 },
 {
  "title": "odyssey",
  "price": 7.5,
  "released": 733622400
 }
]`

func main() {
	// store := list{
	// 	{Title: "moby dick", Price: 10, Released: toTimestamp(118281600)},
	// 	{Title: "odyssey", Price: 15, Released: toTimestamp("733622400")},
	// 	{Title: "hobbit", Price: 25},
	// }

	var store list

	err := json.Unmarshal([]byte(data), &store)
	if err != nil {
		log.Fatal(err)
	}

	store.discount(.5)

	sort.Sort(store)
	fmt.Println(store)

	sort.Sort(sort.Reverse(store))
	fmt.Println(store)

	sort.Sort(store.byReleaseDate())
	fmt.Println(store)

	data, err := json.MarshalIndent(store, "", " ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))
}
