package main

import (
	"fmt"

	s "github.com/inancgumus/prettyslice"
)

/**
 * Created by sofia on 2019-12-27.
 */

/**
 * Source: Udemy Go (Golang) The Complete Bootcamp Course
 */

func main() {
	s.PrintBacking = true

	data := []float64{10, 25, 30, 50}

	// newData := []float64{99, 100}

	// for i := range newData {
	// 	data[i] = newData[i]
	// }

	n := copy(data, []float64{99, 100})
	fmt.Printf("%d elements copied\n", n)

	saved := make([]float64, len(data))
	copy(saved, data)

	s.Show("Probabilities", data)
	s.Show("Saved", saved)

	fmt.Printf("Is it gonna rain? %.f%% chance.\n",
		(data[0]+data[1]+data[2]+data[3])/float64(len(data)),
	)
}
