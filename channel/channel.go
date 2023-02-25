package main

import "sync"
import "os"

import "fmt"

func testNilChannel(operation string) {
	switch operation {
	case "send":
		testNilChannelSend() /* will cause deadlock */
	case "recv":
		testNilChannelRecv() /* will cause deadlock */
	case "close":
		testNilChannelClose() /* will panic */
	}
}

func testNilChannelSend() {
	var ch chan int
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		ch <- 1
		wg.Done()
	}()

	wg.Wait()
}

func testNilChannelRecv() {
	var ch chan int
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		<-ch
		wg.Done()
	}()

	wg.Wait()
}

func testNilChannelClose() {
	var ch chan int
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		close(ch)
		wg.Done()
	}()

	wg.Wait()
}

func testClosedChannel(operation string) {
	switch operation {
	case "send":
		testClosedChannelSend() /* will panic */
	case "recv":
		testClosedChannelRecv() /* no panic, no deadlock, but ok will be false */
	case "close":
		testClosedChannelClose() /* will panic */
	}
}

func testClosedChannelSend() {
	ch := make(chan int)
	close(ch)
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		ch <- 1
		wg.Done()
	}()

	wg.Wait()
}

func testClosedChannelRecv() {
	ch := make(chan int)
	close(ch)
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		v, ok := <-ch
		fmt.Printf("receive from closed channel v:%d ok:%d\n", v, ok)
		wg.Done()
	}()

	wg.Wait()
}

func testClosedChannelClose() {
	ch := make(chan int)
	close(ch)
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		close(ch)
		wg.Done()
	}()

	wg.Wait()
}

func testNormalChannel(operation string) {
	ch := make(chan int)
	fmt.Printf("len of channel is %d\n", len(ch))
	defer close(ch)
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		ch <- 1
		wg.Done()
	}()

	<-ch
	wg.Wait()
}

func testCloseStructEmpty() {
	fmt.Println("closing chan of empty struct")
	var ch = make(chan struct{})
	close(ch)
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("channel [send|recv|close] [nil|closed|normal|sempty]")
	}

	switch os.Args[2] {
	case "nil":
		testNilChannel(os.Args[1])
	case "closed":
		testClosedChannel(os.Args[1])
	case "normal":
		testNormalChannel(os.Args[1])
	case "sempty":
		testCloseStructEmpty()
	}
}
