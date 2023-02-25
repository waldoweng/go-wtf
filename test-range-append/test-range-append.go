package main

import "fmt"

func rangeAppend() {
	v := []int{1, 2, 3}
	for i := range v {
		v = append(v, i)
	}
	fmt.Printf("%v\n", v)
}

func main() {
	rangeAppend()
}
