package main

import "fmt"

func setToNil(c chan int) {
	c = nil
}

func main() {
	var c chan int = nil
	if c == nil {
		fmt.Println("chan is nil where initialized")
	}

	c = make(chan int)
	if c == nil {
		fmt.Println("chan is nil after make")
	}

	setToNil(c)
	if c == nil {
		fmt.Printf("chan is nil after setToNil")
	}

	fmt.Printf("chan is %v\n", c)
}
