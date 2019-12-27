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

	validUser     = "jack"
	validPassword = "1888"

	accessDeniedMsg    = "Access denied for"
	accessGrantedMsg   = "Access granted to"
	invalidPasswordMsg = "Invalid password for"
)

func main() {
	argCount := len(os.Args) - 1

	if argCount != 2 {
		fmt.Println(usage)
		return
	}

	username := os.Args[1]
	password := os.Args[2]

	if username != validUser {
		fmt.Printf("%s %q\n", accessDeniedMsg, username)
	} else if password != validPassword {
		fmt.Printf("%s %q\n", invalidPasswordMsg, password)
	} else {
		fmt.Printf("%s %q\n", accessGrantedMsg, username)
	}
}
