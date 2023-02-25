package main

import "fmt"

type data struct {
	name string
}

func main() {
	var s []int
	fmt.Printf("s len:%d cap:%d\n", len(s), cap(s)) // 0, 0

	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s1 := slice[2:5:7]
	fmt.Printf("s1:%v len(s1):%d cap(s1):%d\n", s1, len(s1), cap(s1)) // [2, 3, 4], 3, 5

	s2 := s1[2:5]
	fmt.Printf("s2:%v len(s2):%d cap(s2):%d\n", s2, len(s2), cap(s2)) // 3, 3

	s2 = append(s2, 100, 200)
	fmt.Printf("s2:%v len(s2):%d cap(s2):%d\n", s2, len(s2), cap(s2))
	s1[2] = 20

	fmt.Println(slice) // 4 = 20
}
