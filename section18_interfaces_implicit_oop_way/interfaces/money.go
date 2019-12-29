package main

import "fmt"

/**
 * Created by sofia on 2019-12-29.
 */

/**
 * Source: Udemy Go (Golang) The Complete Bootcamp Course
 */

type money float64

func (m money) string() string {
	return fmt.Sprintf("$%.2f", m)
}
