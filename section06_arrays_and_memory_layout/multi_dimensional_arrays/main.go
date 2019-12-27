package main

import "fmt"

/**
 * Created by sofia on 2019-12-27.
 */

/**
 * Source: Udemy Go (Golang) The Complete Bootcamp Course
 */

func main() {
	student1 := [3]float64{5, 6, 1}
	student2 := [3]float64{9, 8, 4}

	var sum float64
	for _, v := range student1 {
		sum += v
	}
	for _, v := range student2 {
		sum += v
	}

	const N = float64(len(student1) * 2)

	fmt.Printf("Avg Grade: %g\n", sum/N)

	students := [2][3]float64{
		[3]float64{5, 6, 1},
		[3]float64{9, 8, 4},
	}

	sum = 0

	for i := range students {
		for j := range students[i] {
			sum += students[i][j]
		}
	}

	numStu := float64(len(students) * len(students[0]))

	fmt.Printf("Avg Grade: %g\n", sum/numStu)

	students = [...][3]float64{
		{5, 6, 1},
		{9, 8, 4},
	}

	sum = 0

	for _, grades := range students {
		for _, grade := range grades {
			sum += grade
		}
	}

	fmt.Printf("Avg Grade: %g\n", sum/numStu)
}
