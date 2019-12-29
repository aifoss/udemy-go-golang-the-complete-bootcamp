package main

import (
	"fmt"
	"strings"
)

/**
 * Created by sofia on 2019-12-28.
 */

/**
 * Source: Udemy Go (Golang) The Complete Bootcamp Course
 */

func main() {
	fmt.Println("***** ARRAYS")
	arrays()

	fmt.Println("\n***** SLICES")
	slices()

	fmt.Println("\n***** MAPS")
	maps()

	fmt.Println("\n***** STRUCTS")
	structs()
}

type house struct {
	name  string
	rooms int
}

func structs() {
	myHouse := house{name: "My House", rooms: 5}

	addRoom(myHouse)
	fmt.Printf("%+v\n", myHouse)

	addRoomPtr(&myHouse)
	fmt.Printf("%+v\n", myHouse)
}

func addRoom(h house) {
	h.rooms++
}

func addRoomPtr(h *house) {
	h.rooms++
}

func maps() {
	confused := map[string]int{"one": 2, "two": 1}
	fix(confused)
	fmt.Println(confused)
}

func fix(m map[string]int) {
	m["one"] = 1
	m["two"] = 2
	m["three"] = 3
}

func slices() {
	dirs := []string{"up", "down", "left", "right"}

	up(dirs)
	fmt.Println(dirs)

	upPtr(&dirs)
	fmt.Println(dirs)
}

func up(dirs []string) {
	for i := range dirs {
		dirs[i] = strings.ToUpper(dirs[i])
	}

	dirs = append(dirs, "HEISEN BUG")
}

func upPtr(dirs *[]string) {
	pd := *dirs

	for i := range pd {
		pd[i] = strings.ToUpper(pd[i])
	}

	*dirs = append(*dirs, "HEISEN BUG")
}

func arrays() {
	nums := [...]int{1, 2, 3}

	incr(nums)
	fmt.Println(nums)

	incrByPtr(&nums)
	fmt.Println(nums)
}

func incr(nums [3]int) {
	for i := range nums {
		nums[i]++
	}
}

func incrByPtr(nums *[3]int) {
	for i := range nums {
		nums[i]++
	}
}
