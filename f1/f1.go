package main

import (
	"fmt"

	"github.com/form3tech-oss/f1/v2/pkg/f1"
	"github.com/form3tech-oss/f1/v2/pkg/f1/testing"
)

func main() {
	// Create a new f1 instance, add all the scenarios and execute the f1 tool.
	// Any scenario that is added here can be executed like: `go run main.go run constant mySuperFastLoadTest`
	f1.New().Add("mySuperFastLoadTest", setupMySuperFastLoadTest).Execute()
}

// Performs any setup steps and returns a function to run on every iteration of the scenario
func setupMySuperFastLoadTest(t *testing.T) testing.RunFn {
	fmt.Println("Setup the scenario")

	// Register clean up function which will be invoked at the end of the scenario execution to clean up the setup
	t.Cleanup(func() {
		fmt.Println("Clean up the setup of the scenario")
	})

	runFn := func(t *testing.T) {
		fmt.Println("Run the test")

		// Register clean up function for each test which will be invoked in LIFO order after each iteration
		t.Cleanup(func() {
			fmt.Println("Clean up the test execution")
		})
	}

	return runFn
}
