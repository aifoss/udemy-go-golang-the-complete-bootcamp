package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

/**
 * Created by sofia on 2019-12-27.
 */

/**
 * Source: Udemy Go (Golang) The Complete Bootcamp Course
 */

const usage = "[your name]"

func main() {
	moods := [...]string{
		"sad",
		"terrible",
		"happy",
		"awesome",
		"good",
	}

	args := os.Args[1:]

	if len(args) != 1 {
		fmt.Println(usage)
		return
	}

	name := args[0]

	rand.Seed(time.Now().UnixNano())

	idx := rand.Intn(len(moods))
	mood := moods[idx]

	fmt.Printf("%s feels %s\n", name, mood)
}
