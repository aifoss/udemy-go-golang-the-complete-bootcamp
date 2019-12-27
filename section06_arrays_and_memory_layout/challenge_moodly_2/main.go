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

const usage = "[your name] [positive|negative]"

func main() {
	moods := [...][3]string{
		{"good", "great", "awesome"},
		{"bad", "sad", "terrible"},
	}

	args := os.Args[1:]

	if len(args) != 2 {
		fmt.Println(usage)
		return
	}

	rand.Seed(time.Now().UnixNano())

	idx := rand.Intn(len(moods[0]))

	name := args[0]
	dir := args[1]
	mood := ""

	switch dir {
	case "positive":
		mood = moods[0][idx]
	case "negative":
		mood = moods[1][idx]
	}

	fmt.Printf("%s feels %s\n", name, mood)
}
