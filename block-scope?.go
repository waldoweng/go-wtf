package main

import "log"

func main() {
	log.Println("enter main scope")
	defer func() { log.Println("leaving main scope?") }()

	{
		// in fact if you declare a variable inside the nest scope
		// and use it outside the scope, you get a compile-time error.
		var i int
		log.Println("enter nested scope?")

		// but `defer` seems not bind to this nested-scope.
		defer func() { log.Println("leaving nested scope?") }()
	}

	// uncomment this line you get compile-time error.
	// i = 0
	log.Println("i have leave the nested scope and going to leave the main scope?")
}

// what i expect was:
// 1. enter main scope
// 2. enter nested scope?
// 3. leaving nested scope?
// 4. i have leave the nested scope and going to leave the main scope?
// 5. leaving main scope?

// what i have got was:
// 1. enter main scope
// 2. enter nested scope?
// 3. i have leave the nested scope and going to leave the main scope?
// 4. leaving nested scope?
// 5. leaving main scope?