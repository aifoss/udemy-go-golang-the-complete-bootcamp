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
		parsed := parse(p, in.Text())
		update(p, parsed)
	}

	summarize(p)

	dumpErrs([]error{in.Err(), err(p)})
}

func summarize(p *parser) {
	sort.Strings(p.domains)

	fmt.Printf("%-30s %10s\n", "DOMAIN", "VISITS")
	fmt.Println(strings.Repeat("_", 45))

	for _, domain := range p.domains {
		fmt.Printf("%-30s %10d\n", domain, p.sum[domain].visits)
	}

	fmt.Printf("\n%-30s %10d\n", "TOTAL", p.total)
}

func dumpErrs(errs []error) {
	for _, err := range errs {
		if err != nil {
			fmt.Println("> Err:", err)
		}
	}
}
