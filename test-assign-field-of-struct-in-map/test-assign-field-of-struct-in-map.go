package main

import "fmt"

type data struct {
	name string
}

func main() {
	m := map[string]*data{
		"x": {name: "Tom"},
	}
	m["x"].name = "Jerry"
	fmt.Printf("m: %#v\n", m)
	fmt.Printf("name: %s\n", m["x"].name)
}
