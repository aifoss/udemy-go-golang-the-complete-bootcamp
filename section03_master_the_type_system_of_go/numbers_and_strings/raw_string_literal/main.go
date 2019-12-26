package main

import "fmt"

/**
 * Created by sofia on 2019-12-24.
 */

/**
 * Source: Udemy Go (Golang) The Complete Bootcamp Course
 */

func main() {
	var s string
	s = "how are you?"
	s = `how are you?`

	fmt.Println(s)

	s = "<html>\n\t<body>\"Hello\"</body>\n<html>"

	fmt.Println(s)

	s = `
<html>
	<body>"Hello"</body>
</html>`

	fmt.Println(s)

	fmt.Println("c:\\my\\dir\\file")
	fmt.Println(`c:\my\dir\file`)
}
