package main

import "fmt"

/**
 * Created by sofia on 2019-12-28.
 */

/**
 * Source: Udemy Go (Golang) The Complete Bootcamp Course
 */

func main() {
	var (
		counter int
		V       int
		P       *int
	)

	counter = 100
	P = &counter
	V = *P

	fmt.Printf("counter: %-13d addr: %-13p\n", counter, &counter)
	fmt.Printf("P      : %-13p addr: %-13p *P: %-13d\n", P, &P, *P)
	fmt.Printf("V      : %-13d addr: %-13p\n", V, &V)

	fmt.Println("\n******** change counter")

	counter = 10

	fmt.Printf("counter: %-13d addr: %-13p\n", counter, &counter)
	fmt.Printf("P      : %-13p addr: %-13p *P: %-13d\n", P, &P, *P)
	fmt.Printf("V      : %-13d addr: %-13p\n", V, &V)

	fmt.Println("\n******** change counter in passVal()")

	passVal(counter)

	fmt.Printf("counter: %-13d addr: %-13p\n", counter, &counter)
	fmt.Printf("P      : %-13p addr: %-13p *P: %-13d\n", P, &P, *P)
	fmt.Printf("V      : %-13d addr: %-13p\n", V, &V)

	fmt.Println("\n******** change counter in returnVal()")

	counter = returnVal(counter)

	fmt.Printf("counter: %-13d addr: %-13p\n", counter, &counter)
	fmt.Printf("P      : %-13p addr: %-13p *P: %-13d\n", P, &P, *P)
	fmt.Printf("V      : %-13d addr: %-13p\n", V, &V)

	fmt.Println("\n******** change counter in passPtrVal()")

	passPtrVal(&counter)

	fmt.Printf("counter: %-13d addr: %-13p\n", counter, &counter)
	fmt.Printf("P      : %-13p addr: %-13p *P: %-13d\n", P, &P, *P)
	fmt.Printf("V      : %-13d addr: %-13p\n", V, &V)
}

func passVal(n int) {
	n = 50
	fmt.Printf("n      : %-13d addr: %-13p\n", n, &n)
}

func returnVal(r int) int {
	r = 50
	fmt.Printf("r      : %-13d addr: %-13p\n", r, &r)
	return r
}

func passPtrVal(pn *int) {
	fmt.Printf("pn     : %-13p addr: %-13p *P: %-13d\n", pn, &pn, *pn)

	*pn++

	fmt.Printf("pn     : %-13p addr: %-13p *P: %-13d\n", pn, &pn, *pn)
}
