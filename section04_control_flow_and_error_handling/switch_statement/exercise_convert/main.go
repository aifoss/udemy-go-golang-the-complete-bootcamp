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

// ---------------------------------------------------------
// EXERCISE: Convert
//
//  Convert the if statement to a switch statement.
// ---------------------------------------------------------

const (
	usage       = "Usage: [username] [password]"
	errUser     = "Access denied for %q.\n"
	errPwd      = "Invalid password for %q.\n"
	accessOK    = "Access granted to %q.\n"
	user, user2 = "jack", "inanc"
	pass, pass2 = "1888", "1879"
)

func main() {
	args := os.Args

	if len(args) != 3 {
		fmt.Println(usage)
		return
	}

	u, p := args[1], args[2]

	//
	// REFACTOR THIS TO A SWITCH
	//
	if u != user && u != user2 {
		fmt.Printf(errUser, u)
	} else if u == user && p == pass {
		fmt.Printf(accessOK, u)
	} else if u == user2 && p == pass2 {
		fmt.Printf(accessOK, u)
	} else {
		fmt.Printf(errPwd, u)
	}

	switch u {
	case user:
		switch p {
		case pass:
			fmt.Printf(accessOK, u)
		default:
			fmt.Printf(errPwd, u)
		}
	case user2:
		switch p {
		case pass2:
			fmt.Printf(accessOK, u)
		default:
			fmt.Printf(errPwd, u)
		}
	default:
		fmt.Printf(errUser, u)
	}

	switch {
	case u != user && u != user2:
		fmt.Printf(errUser, u)
	case u == user && p == pass:
		fmt.Printf(accessOK, u)
	case u == user2 && p == pass2:
		fmt.Printf(accessOK, u)
	default:
		fmt.Printf(errPwd, u)
	}
}
