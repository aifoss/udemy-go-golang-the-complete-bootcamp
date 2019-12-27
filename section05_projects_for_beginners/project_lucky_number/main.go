package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

/**
 * Created by sofia on 2019-12-26.
 */

/**
 * Source: Udemy Go (Golang) The Complete Bootcamp Course
 */

// ---------------------------------------------------------
// EXERCISE: First Turn Winner
//
//  If the player wins on the first turn, then display
//  a special bonus message.
//
// RESTRICTION
//  The first picked random number by the computer should
//  match the player's guess.
//
// EXAMPLE SCENARIO
//  1. The player guesses 6
//  2. The computer picks a random number and it happens
//     to be 6
//  3. Terminate the game and print the bonus message
// ---------------------------------------------------------

// ---------------------------------------------------------
// EXERCISE: Random Messages
//
//  Display a few different won and lost messages "randomly".
//
// HINTS
//  1. You can use a switch statement to do that.
//  2. Set its condition to the random number generator.
//  3. I would use a short switch.
//
// EXAMPLES
//  The Player wins: then randomly print one of these:
//
//  go run main.go 5
//    YOU WON
//  go run main.go 5
//    YOU'RE AWESOME
//
//  The Player loses: then randomly print one of these:
//
//  go run main.go 5
//    LOSER!
//  go run main.go 5
//    YOU LOST. TRY AGAIN?
// ---------------------------------------------------------

// ---------------------------------------------------------
// EXERCISE: Double Guesses
//
//  Let the player guess two numbers instead of one.
//
// HINT:
//  Generate random numbers using the greatest number
//  among the guessed numbers.
//
// EXAMPLES
//  go run main.go 5 6
//  Player wins if the random number is either 5 or 6.
// ---------------------------------------------------------

// ---------------------------------------------------------
// EXERCISE: Enough Picks
//
//  If the player's guess number is below 10;
//  then make sure that the computer generates a random
//  number between 0 and 10.
//
//  However, if the guess number is above 10; then the
//  computer should generate the numbers
//  between 0 and the guess number.
//
// WHY?
//  This way the game will be harder.
//
//  Because, in the current version of the game, if
//  the player's guess number is for example 3; then the
//  computer picks a random number between 0 and 3.
//
//  So, currently a player can easily win the game.
//
// EXAMPLE
//  Suppose that the player runs the game like this:
//    go run main.go 9
//
//  Or like this:
//    go run main.go 2
//
//    Then the computer should pick a random number
//    between 0-10.
//
//  Or, if the player runs it like this:
//    go run main.go 15
//
//    Then the computer should pick a random number
//    between 0-15.
// ---------------------------------------------------------

// ---------------------------------------------------------
// EXERCISE: Dynamic Difficulty
//
//  Current game picks only 5 numbers (5 turns).
//
//  Make sure that the game adjust its own difficulty
//  depending on the guess number.
//
// RESTRICTION
//  Do not make the game to easy. Only adjust the
//  difficulty if the guess is above 10.
//
// EXPECTED OUTPUT
//  Suppose that the player runs the game like this:
//    go run main.go 5
//
//    Then the computer should pick 5 random numbers.
//
//  Or, if the player runs it like this:
//    go run main.go 25
//
//    Then the computer may pick 11 random numbers
//    instead.
//
//  Or, if the player runs it like this:
//    go run main.go 100
//
//    Then the computer may pick 30 random numbers
//    instead.
//
//  As you can see, greater guess number causes the
//  game to increase the game turns, which in turn
//  adjust the game's difficulty dynamically.
//
//  Because, greater guess number makes it harder to win.
//  But, this way, game's difficulty will be dynamic.
//  It will adjust its own difficulty depending on the
//  guess number.
// ---------------------------------------------------------

const (
	minMax = 10
	usage  = `Welcome to the Lucky Number Game!

The program will pick random numbers.
Your mission is to guess one of those numbers

The greater your number is, the harder it gets.

Wanna play?
`
	bonusMsg = "WOW! You won on the first try! You ROCK!!!"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	args := os.Args[1:]

	if len(args) < 1 {
		fmt.Printf(usage)
		return
	}

	guess1, err := strconv.Atoi(args[0])

	if err != nil {
		fmt.Println("Not a number")
		return
	}

	var guess2 int

	if len(args) == 2 {
		guess2, err = strconv.Atoi(args[1])

		if err != nil {
			fmt.Println("Not a number")
			return
		}
	}

	if guess1 < 0 || guess2 < 0 {
		fmt.Println("Please pick a positive number")
		return
	}

	max := guess1

	if guess2 > guess1 {
		max = guess2
	}

	if max < minMax {
		max = minMax
	}

	maxTurns := 5

	switch {
	case max > 100:
		maxTurns = 30
	case max > 50:
		maxTurns = 20
	case max > 10:
		maxTurns = 10
	}

	fmt.Printf("We will play the game up to %d turns\n", maxTurns)

	winMsgs := []string{"YOU WIN!", "You are the winner!", "You guessed right :)"}
	loseMsgs := []string{"YOU LOST!", "You lost, try again?", "You guessed wrong :("}

	for turn := 1; turn <= maxTurns; turn++ {
		n := rand.Intn(max + 1)

		fmt.Printf("Computer picked %d\n", n)

		if n == guess1 || n == guess2 {
			winMsgIdx := rand.Intn(len(winMsgs))

			fmt.Println(winMsgs[winMsgIdx])
			fmt.Printf("Winning number is %d\n", n)

			if turn == 1 {
				fmt.Println(bonusMsg)
			}

			return
		}
	}

	loseMsgIdx := rand.Intn(len(loseMsgs))

	fmt.Println(loseMsgs[loseMsgIdx])
}
