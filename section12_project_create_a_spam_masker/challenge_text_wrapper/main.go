package main

import (
	"fmt"
	"unicode"
)

/**
 * Created by sofia on 2019-12-28.
 */

/**
 * Source: Udemy Go (Golang) The Complete Bootcamp Course
 */

// ---------------------------------------------------------
// In this example your goal is to mimic the soft text wrapping feature of text editors.
// For example, when there are 100 characters on a line and if the soft-wrapping is set to 40,
// an editor may cut the line that goes beyond 40 characters and display the rest of the line
// in the next line instead.
//
// EXAMPLE
//
// Wrap the given text for 40 characters per line.
// For example, for the following input, the program should print the following output.
//
// INPUT:
//
// Hello world, how is it going? It is ok.. The weather is beautiful.
//
// OUTPUT:
//
// Hello world, how is it going? It is ok..
// The weather is beautiful.
//
// RULES
//
// The program should work with Unicode text. You can find a unicode text in story.txt file.
//
// The program should not cut the words before they finish. Instead, it should put the whole word on the next line. For example, this is not OK:
//
// Hello world, how is it goi
// ng? It is o
// k. The weather is beautifu
// l.
// ---------------------------------------------------------

func main() {
	const text = `Galaksinin Batı Sarmal Kolu'nun bir ucunda, haritası bile çıkarılmamış ücra bir köşede, gözlerden uzak, küçük ve sarı bir güneş vardır.

Bu güneşin yörüngesinde, kabaca yüz kırksekiz milyon kilometre uzağında, tamamıyla önemsiz ve mavi-yeşil renkli, küçük bir gezegen döner.

Gezegenin maymun soyundan gelen canlıları öyle ilkeldir ki dijital kol saatinin hâlâ çok etkileyici bir buluş olduğunu düşünürler.`

	const maxWidth = 40

	var lw int // line width

	for _, r := range text {
		fmt.Printf("%c", r)

		switch lw++; {
		case lw > maxWidth && r != '\n' && unicode.IsSpace(r):
			fmt.Println()
			fallthrough
		case r == '\n':
			lw = 0
		}
	}

}
