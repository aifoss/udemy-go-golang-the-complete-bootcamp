package main

import (
	"fmt"
	"unsafe"

	s "github.com/inancgumus/prettyslice"
)

/**
 * Created by sofia on 2019-12-27.
 */

/**
 * Source: Udemy Go (Golang) The Complete Bootcamp Course
 */

type collection [4]string
type collection2 []string

func main() {
	s.PrintElementAddr = true

	data := collection{"slices", "are", "awesome", "period"}
	change(data)

	s.Show("main's data", data)
	fmt.Printf("main's data array address: %p\n", &data)

	data2 := collection2{"slices", "are", "awesome", "period"}
	change2(data2)

	s.Show("main's data2", data2)
	fmt.Printf("main's data2 slice address: %p\n", &data2)

	fmt.Printf("array's size: %d bytes\n", unsafe.Sizeof(data))
	fmt.Printf("slice's size: %d bytes\n", unsafe.Sizeof(data2))
}

func change(data collection) {
	data[2] = "brilliant!"
	s.Show("change's data", data)
	fmt.Printf("change's data array address: %p\n", &data)
}

func change2(data2 collection2) {
	data2[2] = "brilliant!"
	s.Show("change's data2", data2)
	fmt.Printf("change's data2 slice address: %p\n", &data2)
}
