package main

import "fmt"

/**
 * Created by sofia on 2019-12-28.
 */

/**
 * Source: Udemy Go (Golang) The Complete Bootcamp Course
 */

// func main() {
// 	var schools []map[int]string

// 	schools = append(schools, make(map[int]string))
// 	load(&schools[0], &([]string{"batman", "superman"}))

// 	schools = append(schools, make(map[int]string))
// 	load(&schools[1], &([]string{"spiderman", "wonder woman"}))

// 	fmt.Println(schools[0])
// 	fmt.Println(schools[1])
// }

// func load(m *map[int]string, students *[]string) {
// 	for i, name := range *students {
// 		(*m)[i+1] = name
// 	}
// }

func main() {
	schools := make([]map[int]string, 2)
	for i := range schools {
		schools[i] = make(map[int]string)
	}

	load(schools[0], []string{"batman", "superman"})
	load(schools[1], []string{"spiderman", "wonder woman"})

	fmt.Println(schools[0])
	fmt.Println(schools[1])
}

func load(m map[int]string, students []string) {
	for i, name := range students {
		m[i+1] = name
	}
}
