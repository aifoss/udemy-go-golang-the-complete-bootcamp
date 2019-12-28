package main

import (
	"fmt"
	"unsafe"
)

/**
 * Created by sofia on 2019-12-28.
 */

/**
 * Source: Udemy Go (Golang) The Complete Bootcamp Course
 */

func main() {
	empty := ""
	dump(empty)

	hello := "hello"
	dump(hello)
	dump("hello")

	for i := range hello {
		dump(hello[i : i+1])
	}

	dump(string([]byte(hello)))
	dump(string([]byte(hello)))
	dump(string([]rune(hello)))
}

// StringHeader is used by a string value
type StringHeader struct {
	pointer uintptr // where it starts
	length  int     // where it ends
}

// dump prints the stinrg header of a string value
func dump(s string) {
	ptr := *(*StringHeader)(unsafe.Pointer(&s))
	fmt.Printf("%q: %+v\n", s, ptr)
}
