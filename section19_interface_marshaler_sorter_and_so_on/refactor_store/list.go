package main

import (
	"fmt"
	"sort"
	"strings"
)

/**
 * Created by sofia on 2019-12-29.
 */

/**
 * Source: Udemy Go (Golang) The Complete Bootcamp Course
 */

type list []*product

func (l list) String() string {
	if len(l) == 0 {
		return fmt.Sprintf("Sorry. We're waiting for delivery.\n")
	}

	var str strings.Builder

	for _, p := range l {
		str.WriteString("* ")
		str.WriteString(p.String())
		str.WriteRune('\n')
	}

	return str.String()
}

func (l list) discount(ratio float64) {
	for _, p := range l {
		p.discount(ratio)
	}
}

func (l list) Len() int {
	return len(l)
}

func (l list) Less(i, j int) bool {
	return l[i].Title < l[j].Title
}

func (l list) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

type byRelease struct {
	list
}

func (br byRelease) Less(i, j int) bool {
	return br.list[i].Released.Before(br.list[j].Released.Time)
}

func (l list) byReleaseDate() sort.Interface {
	return &byRelease{l}
}
