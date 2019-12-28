package main

import (
	"fmt"
	"path"
)

/**
 * Created by sofia on 2019-12-23.
 */

/**
 * Source: Udemy Go (Golang) The Complete Bootcamp Course
 */

func main() {
	var dir, file string

	dir, file = path.Split("css/main.css")

	fmt.Println(dir, file)

	dir, _ = path.Split("src/package/main.go")

	fmt.Println("dir:", dir)

	_, file = path.Split("src/package/main.go")

	fmt.Println("file:", file)

	d, f := path.Split("css/main.css")

	fmt.Println(d, f)
}
