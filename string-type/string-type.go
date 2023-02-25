package main

import "fmt"

func main() {
	var s string = "string1"
	fmt.Println(s)

	s = "string"
	fmt.Println(s)

	// won't compile
	// s = nil
}
