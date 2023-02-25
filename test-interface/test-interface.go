package main

import "fmt"
import "reflect"

func main() {
	err := run()
	if reflect.ValueOf(err).IsNil() {
		fmt.Println("err is nil")
	} else {
		fmt.Println(err.Error())
	}
}

func run() (err error) {
	return check()
}

func check() *Result {
	var res *Result
	return res
}

type Result struct {
	message string
}

func (result *Result) Error() string {
	return result.message
}
