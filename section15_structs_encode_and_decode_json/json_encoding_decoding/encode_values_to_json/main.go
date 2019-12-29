package main

import (
	"encoding/json"
	"fmt"
)

/**
 * Created by sofia on 2019-12-28.
 */

/**
 * Source: Udemy Go (Golang) The Complete Bootcamp Course
 */

type permissions map[string]bool

type user struct {
	Name        string      `json:"username"`
	Password    string      `json:"-"`
	Permissions permissions `json:"perms,omitempty"`
}

func main() {
	users := []user{
		{"inanc", "1234", nil},
		{"john", "42", permissions{"admin": true}},
		{"tom", "666", permissions{"write": true}},
	}

	out, err := json.MarshalIndent(users, "", "\t")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(out))
}
