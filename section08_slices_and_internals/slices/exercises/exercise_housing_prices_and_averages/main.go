package main

import (
	"fmt"
	"strconv"
	"strings"
)

/**
 * Created by sofia on 2019-12-27.
 */

/**
 * Source: Udemy Go (Golang) The Complete Bootcamp Course
 */

// ---------------------------------------------------------
// EXERCISE: Housing Prices and Averages
//
//  Use the previous exercise to solve this exercise (Housing Prices).
//
//  Your task is to print the averages of the sizes, beds, baths, and prices.
//
//
// EXPECTED OUTPUT
//
//  Location       Size           Beds           Baths          Price
//  ===========================================================================
//  New York       100            2              1              100000
//  New York       150            3              2              200000
//  Paris          200            4              3              400000
//  Istanbul       500            10             5              1000000
//
//  ===========================================================================
//                 237.50         4.75           2.75           425000.00
//
// ---------------------------------------------------------

func main() {
	const (
		header = "Location,Size,Beds,Baths,Price"
		data   = `New York,100,2,1,100000
New York,150,3,2,200000
Paris,200,4,3,400000
Istanbul,500,10,5,1000000`
		separator = ","
		newline   = "\n"
	)

	var (
		locations []string
		sizes     []int
		beds      []int
		baths     []int
		prices    []int
	)

	var headers []string = strings.Split(header, separator)
	var rows []string = strings.Split(data, newline)

	for _, row := range rows {
		vals := strings.Split(row, separator)

		for i := range vals {
			if i == 0 {
				locations = append(locations, vals[i])
			} else {
				val, _ := strconv.Atoi(vals[i])

				switch i {
				case 1:
					sizes = append(sizes, val)
				case 2:
					beds = append(beds, val)
				case 3:
					baths = append(baths, val)
				case 4:
					prices = append(prices, val)
				}
			}
		}
	}

	numData := len(rows)
	var averages []float64

	for i := 1; i < len(headers); i++ {
		sum := 0.0

		switch i {
		case 1:
			for _, v := range sizes {
				sum += float64(v)
			}
		case 2:
			for _, v := range beds {
				sum += float64(v)
			}
		case 3:
			for _, v := range baths {
				sum += float64(v)
			}
		case 4:
			for _, v := range prices {
				sum += float64(v)
			}
		}

		averages = append(averages, sum/float64(numData))
	}

	for _, h := range headers {
		fmt.Printf("%-15s", h)
	}
	fmt.Println()

	fmt.Println(strings.Repeat("-", 70))

	for i := 0; i <= numData+3; i++ {
		if i < numData {
			fmt.Printf("%-15s", locations[i])
			fmt.Printf("%-15d", sizes[i])
			fmt.Printf("%-15d", beds[i])
			fmt.Printf("%-15d", baths[i])
			fmt.Printf("%-15d", prices[i])
			fmt.Println()
		} else if i == numData+1 {
			fmt.Println()
		} else if i == numData+2 {
			fmt.Println(strings.Repeat("-", 70))
		} else if i == numData+3 {
			fmt.Printf("%-15s", " ")
			for j := range averages {
				fmt.Printf("%-15.2f", averages[j])
			}
			fmt.Println()
		}
	}
}
