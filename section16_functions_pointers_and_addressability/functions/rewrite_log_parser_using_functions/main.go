package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

/**
 * Created by sofia on 2019-12-28.
 */

/**
 * Source: Udemy Go (Golang) The Complete Bootcamp Course
 */

func main() {
	p := newParser()

	in := bufio.NewScanner(os.Stdin)

	for in.Scan() {
		p.lines++

		parsed, err := parse(p, in.Text())
		if err != nil {
			fmt.Println(err)
			return
		}

		p = update(p, parsed)
	}

	sort.Strings(p.domains)

	fmt.Printf("%-30s %10s\n", "DOMAIN", "VISITS")
	fmt.Println(strings.Repeat("_", 45))

	for _, domain := range p.domains {
		parsed := p.sum[domain]
		fmt.Printf("%-30s %10d\n", domain, parsed.visits)
	}

	fmt.Printf("\n%-30s %10d\n", "TOTAL", p.total)

	if err := in.Err(); err != nil {
		fmt.Println("> Err:", err)
	}
}
