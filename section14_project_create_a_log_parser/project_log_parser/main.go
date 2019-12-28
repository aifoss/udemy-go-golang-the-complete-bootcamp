package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

/**
 * Created by sofia on 2019-12-28.
 */

/**
 * Source: Udemy Go (Golang) The Complete Bootcamp Course
 */

func main() {
	var (
		sum     map[string]int // total visits per domain
		domains []string       // unique domain names
		total   int            // total visits to all domains
		lines   int            // number of parsed lines
	)

	sum = make(map[string]int)

	in := bufio.NewScanner(os.Stdin)

	for in.Scan() {
		lines++

		fields := strings.Fields(in.Text())
		if len(fields) != 2 {
			fmt.Printf("wrong input: %v (line #%d)\n", fields, lines)
			return
		}

		domain := fields[0]

		visits, err := strconv.Atoi(fields[1])
		if visits < 0 || err != nil {
			fmt.Printf("wrong input: %q (line #%d)\n", fields[1], lines)
			return
		}

		if _, ok := sum[domain]; !ok {
			domains = append(domains, domain)
		}

		total += visits
		sum[domain] += visits
	}

	sort.Strings(domains)

	fmt.Printf("%-30s %10s\n", "DOMAIN", "VISITS")
	fmt.Println(strings.Repeat("_", 45))

	for _, domain := range domains {
		visits := sum[domain]
		fmt.Printf("%-30s %10d\n", domain, visits)
	}

	fmt.Printf("\n%-30s %10d\n", "TOTAL", total)

	if err := in.Err(); err != nil {
		fmt.Println("> Err:", err)
	}
}
