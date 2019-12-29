package main

import (
	"fmt"
	"strconv"
)

/**
 * Created by sofia on 2019-12-28.
 */

/**
 * Source: Udemy Go (Golang) The Complete Bootcamp Course
 */

func main() {
	local := 10
	show(local)

	local = incr(local)
	show(local)

	local = multBy(local, 5)
	show(local)

	_, err := multByStr(local, "TWO")
	if err != nil {
		fmt.Printf("err -> %s\n", err)
	}

	local, _ = multByStr(local, "2")
	show(local)

	local = sanitize(multByStr(local, "2"))
	show(local)

	local = multBy(local, 5)
	show(local)

	local = limit(local, 500)
	show(local)
}

func show(n int) {
	fmt.Printf("show -> n = %d\n", n)
}

func incr(n int) int {
	n++
	return n
}

func multBy(n, factor int) int {
	return n * factor
}

func multByStr(n int, factor string) (int, error) {
	m, err := strconv.Atoi(factor)
	// if err == nil {
	n = multBy(n, m)
	// }
	return n, err
}

func sanitize(n int, err error) int {
	if err != nil {
		return 0
	}
	return n
}

func limit(n, lim int) (m int) {
	m = n
	if m >= lim {
		m = lim
	}
	return
}
