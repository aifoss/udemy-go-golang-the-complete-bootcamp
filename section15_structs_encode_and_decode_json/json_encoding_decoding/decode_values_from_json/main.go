package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

/**
 * Created by sofia on 2019-12-28.
 */

/**
 * Source: Udemy Go (Golang) The Complete Bootcamp Course
 */

type user struct {
	Name        string          `json:"username"`
	Permissions map[string]bool `json:"perms"`
}

func main() {
	var input []byte

	for in := bufio.NewScanner(os.Stdin); in.Scan(); {
		input = append(input, in.Bytes()...)
	}

	var users []user

	err := json.Unmarshal(input, &users)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(users)
	fmt.Println()

	for _, user := range users {
		fmt.Print("+ ", user.Name)

		switch p := user.Permissions; {
		case p == nil:
			fmt.Print(" has no power")
		case p["admin"]:
			fmt.Print(" is an admin")
		case p["write"]:
			fmt.Print(" can write")
		}

		fmt.Println()
	}
}
