package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/**
 * Created by sofia on 2019-12-28.
 */

/**
 * Source: Udemy Go (Golang) The Complete Bootcamp Course
 */

// ---------------------------------------------------------
// EXERCISE: Refactor the game store to funcs - step #1
//
//  Remember the game store program from the structs exercises?
//  Now, it's time to refactor it to funcs.
//
//  Create games.go file
//
//     1- Move the structs there
//
//     2- Add a func that creates and returns a game.
//
//        Name  : newGame
//        Inputs: id, price, name and genre
//        Output: game
//
//     3- Add a func that adds a `game` to `[]game` and returns `[]game`:
//
//        Name  : addGame
//        Inputs: []game, game
//        Output: []game
//
//     4- Add a func that uses newGame and addGame funcs to load up the game
//        store:
//
//        Name  : load
//        Inputs: none
//        Output: []game
//
//     5- Add a func that indexes games by ID:
//
//        Name  : indexByID
//        Inputs: []game
//        Output: map[int]game
//
//     6- Add a func that prints the given game:
//
//        Name  : printGame
//        Inputs: game
//
//
//  Go back to main.go and change the existing code with
//  the new funcs that you've created. There are hints
//  inside the code.
//
// ---------------------------------------------------------

// ---------------------------------------------------------
// EXERCISE: Refactor the game store to funcs - step #2
//
//  Let's continue the refactoring from the previous
//  exercise. This time, you're going to refactor the
//  command handling logic.
//
//
//  Create commands.go file
//
//     1- Add a func that runs the given command from the user:
//
//        Name  : runCmd
//        Inputs: input string, []game, map[int]game
//        Output: bool
//
//        This func returns true if it wants the program to
//        continue. When it returns false, the program will
//        terminate. So, all the commands that it calls need
//        to return true or false as well depending on whether
//        they want to cause the program to terminate or not.
//
//     2- Add a func that handles the quit command:
//
//        Name  : cmdQuit
//        Input : none
//        Output: bool
//
//     3- Add a func that handles the list command:
//
//        Name  : cmdList
//        Inputs: []game
//        Output: bool
//
//     4- Add a func that handles the id command:
//
//        Name  : cmdByID
//        Inputs: cmd []string, []game, map[int]game
//        Output: bool
//
//     5- Refactor the runCmd() with the cmdXXX funcs.
//
//  Go back to main.go and change the existing code with
//  the new funcs that you've created. There are hints
//  inside the code.
//
// ---------------------------------------------------------

// ---------------------------------------------------------
// EXERCISE: Refactor the game store to funcs - step #3
//
//  Let's continue from the previous exercise. This time,
//  you're going to add json encoding and decoding using
//  funcs.
//
//  1- Create a new command in a func that encodes the
//     game store data to json and terminates the program.
//     Lastly, add it to runCmd func.
//
//     For more information, see: "Encode" exercise from
//     the structs section.
//
//        Name  : cmdSave
//        Inputs: []game
//        Output: bool
//
//  2- Refactor the load() to load the games from the
//     `data` constant (it's in the games.go as well).
//
//     For more information, see: "Decode" exercise from
//     the structs section.
//
// ---------------------------------------------------------

func main() {
	games := load()

	byID := indexByID(games)

	in := bufio.NewScanner(os.Stdin)

	for {
		fmt.Printf(`
  > list   : lists all the games
  > id N   : queries a game by id
  > save   : exports the data to json and quits
  > quit   : quits
`)

		if !in.Scan() {
			break
		}

		fmt.Println()

		cmd := strings.Fields(in.Text())

		if len(cmd) == 0 {
			continue
		}

		cont := runCmd(in.Text(), games, byID)

		if !cont {
			break
		}
	}
}
