package main

import (
	"fmt"
	"time"

	"github.com/inancgumus/screen"
)

/**
 * Created by sofia on 2019-12-27.
 */

/**
 * Source: Udemy Go (Golang) The Complete Bootcamp Course
 */

func main() {

	for {
		screen.MoveTopLeft()

		now := time.Now()
		hour, min, sec := now.Hour(), now.Minute(), now.Second()
		ssec := now.Nanosecond() / 1e8

		clock := [...]placeholder{
			digits[hour/10], digits[hour%10],
			colon,
			digits[min/10], digits[min%10],
			colon,
			digits[sec/10], digits[sec%10],
			dot,
			digits[ssec],
		}

		alarmed := sec%10 == 0

		for line := range clock[0] {
			if alarmed {
				clock = alarm
			}

			for index, digit := range clock {
				next := clock[index][line]
				if (digit == colon || digit == dot) && sec%2 == 0 {
					next = "   "
				}
				fmt.Print(next, "  ")
			}
			fmt.Println()
		}
		fmt.Println()

		const splitSecond = time.Second / 10
		time.Sleep(splitSecond)
	}
}
