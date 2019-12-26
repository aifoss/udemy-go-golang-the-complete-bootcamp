package main

import (
	"fmt"

	"github.com/aifoss/udemy-go-golang-the-complete-bootcamp/section03_master_the_type_system_of_go/go_type_system_mechanics/underlying_types/weights"
)

/**
 * Created by sofia on 2019-12-25.
 */

/**
 * Source: Udemy Go (Golang) The Complete Bootcamp Course
 */

func main() {
	type (
		Gram     int
		Kilogram Gram
		Ton      Kilogram
	)

	var (
		salt   Gram     = 100
		apples Kilogram = 5
		truck  Ton      = 10
	)

	fmt.Printf("salt: %d, apples: %d, truck: %d\n", salt, apples, truck)

	salt = Gram(apples)
	apples = Kilogram(truck)
	truck = Ton(Gram(int(salt)))

	fmt.Printf("salt: %d, apples: %d, truck: %d\n", salt, apples, truck)

	salt = Gram(weights.Gram(100))

	fmt.Printf("Type of weights.Gram: %T\n", weights.Gram(1))
	fmt.Printf("Type of main.Gram: %T\n", Gram(1))

	type myGram weights.Gram

	var pepper myGram = 20

	fmt.Printf("Type of pepper: %T\n", pepper)
}
