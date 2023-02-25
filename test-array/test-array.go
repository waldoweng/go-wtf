package main

import "fmt"

func arrayTest() {
	nums := [5]int{}

	func(a [5]int) {
		for i := 0; i < len(a); i++ {
			a[i] = i
		}
	}(nums)

	fmt.Printf("nums: %+v\n", nums) // [0, 0, 0, 0, 0]

	func(b []int) {
		for i := 0; i < len(b); i++ {
			b[i] = i
		}
	}(nums[2:])

	fmt.Printf("nums: %+v\n", nums) // [0, 0, 0, 1, 2]
}

func main() {
	arrayTest()
}
