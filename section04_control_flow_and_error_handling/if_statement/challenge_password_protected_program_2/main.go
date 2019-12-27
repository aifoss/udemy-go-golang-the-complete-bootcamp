package main

import (
	"fmt"
	"os"
)

/**
 * Created by sofia on 2019-12-26.
 */

/**
 * Source: Udemy Go (Golang) The Complete Bootcamp Course
 */

const (
	usage = "Usage: [username] [password]"

	user1, pass1 = "jack", "1888"
	user2, pass2 = "inanc", "1879"

	errUser  = "Access denied for %q\n"
	errPass  = "Invalid password for %q\n"
	accessOk = "Access granted to %q\n"
)

func main() {
	args := os.Args

	argCount := len(args) - 1

	if argCount != 2 {
		fmt.Println(usage)
		return
	}

	user, pass := args[1], args[2]

	if user != user1 && user != user2 {
		fmt.Printf(errUser, user)
	} else if user == user1 && pass != pass1 {
		fmt.Printf(errPass, user)
	} else if user == user2 && pass != pass2 {
		fmt.Printf(errPass, user)
	} else {
		fmt.Printf(accessOk, user)
	}

}
