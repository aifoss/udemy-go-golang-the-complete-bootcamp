package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

/**
 * Created by sofia on 2019-12-28.
 */

/**
 * Source: Udemy Go (Golang) The Complete Bootcamp Course
 */

func runCmd(input string, games []game, byID map[int]game) bool {
	cmd := strings.Fields(input)

	switch cmd[0] {
	case "quit":
		return cmdQuit()
	case "list":
		return cmdList(games)
	case "id":
		return cmdByID(cmd, games, byID)
	case "save":
		return cmdSave(games)
	default:
		return false
	}
}

func cmdQuit() bool {
	fmt.Println("bye!")
	return false
}

func cmdList(games []game) bool {
	for _, g := range games {
		printGame(g)
	}
	return true
}

func cmdByID(cmd []string, games []game, byID map[int]game) bool {
	if len(cmd) != 2 {
		fmt.Println("wrong id")
		return true
	}

	id, err := strconv.Atoi(cmd[1])
	if err != nil {
		fmt.Println("wrong id")
		return true
	}

	g, ok := byID[id]
	if !ok {
		fmt.Println("sorry. i don't have the game")
		return true
	}

	printGame(g)

	return true
}

func cmdSave(games []game) bool {
	var encodable []jsonGame
	for _, g := range games {
		encodable = append(encodable,
			jsonGame{g.id, g.name, g.genre, g.price})
	}

	out, err := json.MarshalIndent(encodable, "", "\t")
	if err != nil {
		fmt.Println("Sorry:", err)
		return true
	}

	fmt.Println(string(out))

	return false
}
