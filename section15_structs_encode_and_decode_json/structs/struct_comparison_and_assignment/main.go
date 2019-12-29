package main

import "fmt"

/**
 * Created by sofia on 2019-12-28.
 */

/**
 * Source: Udemy Go (Golang) The Complete Bootcamp Course
 */

type song struct {
	title, artist string
}

type playlist struct {
	genre string
	songs []song
}

func main() {
	song1 := song{title: "wonderwall", artist: "oasis"}
	song2 := song{title: "super sonic", artist: "oasis"}

	if song1 == song2 {
		fmt.Println("songs are equal")
	} else {
		fmt.Println("songs are not equal")
	}

	// song1 = song2

	// if song1 == song2 {
	// 	fmt.Println("songs are equal")
	// } else {
	// 	fmt.Println("songs are not equal")
	// }

	fmt.Println()

	rock := playlist{
		genre: "indie rock",
		songs: []song{song1, song2},
	}

	song := rock.songs[0]
	song.title = "live forever"

	fmt.Printf("%+v\n%+v\n\n", song, rock.songs[0])

	rock.songs[0].title = "live forever"

	fmt.Printf("%-20s %20s\n", "TITLE", "ARTIST")
	for _, s := range rock.songs {
		fmt.Printf("%-20s %20s\n", s.title, s.artist)
	}
}
