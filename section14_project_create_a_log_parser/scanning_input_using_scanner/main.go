package main

import (
	"bufio"
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
	// os.Stdin.Close()

	in := bufio.NewScanner(os.Stdin)

	var lines int

	for in.Scan() {
		lines++
		fmt.Println("Scanned text:", in.Text())
	}

	fmt.Printf("%d lines\n", lines)

	// if err := in.Err(); err != nil {
	// 	fmt.Println("ERROR:", err)
	// }
}
