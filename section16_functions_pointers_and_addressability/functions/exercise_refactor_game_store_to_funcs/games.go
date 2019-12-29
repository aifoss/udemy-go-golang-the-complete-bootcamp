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

type item struct {
	id    int
	name  string
	price int
}

type game struct {
	item
	genre string
}

type jsonGame struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Genre string `json:"genre"`
	Price int    `json:"price"`
}

const data = `
[
        {
                "id": 1,
                "name": "god of war",
                "genre": "action adventure",
                "price": 50
        },
        {
                "id": 2,
                "name": "x-com 2",
                "genre": "strategy",
                "price": 40
        },
        {
                "id": 3,
                "name": "minecraft",
                "genre": "sandbox",
                "price": 20
        }
]`

func newGame(id, price int, name, genre string) game {
	return game{
		item:  item{id: id, name: name, price: price},
		genre: genre,
	}
}

func addGame(games []game, g game) []game {
	games = append(games, g)
	return games
}

func load() []game {
	var decoded []jsonGame
	if err := json.Unmarshal([]byte(data), &decoded); err != nil {
		fmt.Println("Sorry, there is a problem:", err)
		return nil
	}

	var games []game
	for _, dg := range decoded {
		games = append(games, game{item{dg.ID, dg.Name, dg.Price}, dg.Genre})
	}

	return games
}

func indexByID(games []game) map[int]game {
	byID := make(map[int]game)
	for _, g := range games {
		byID[g.id] = g
	}
	return byID
}

func printGame(g game) {
	fmt.Printf("#%d: %-15q %-20s $%d\n",
		g.id, g.name, "("+g.genre+")", g.price,
	)
}
