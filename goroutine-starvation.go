package main

import "sync"
import "sync/atomic"
import "log"
import "runtime"
import "time"

func main() {
	// set GOMAXPROCS to 1
	// so that all go-routine will be scheduled to run on this single thread
	runtime.GOMAXPROCS(1)

	// going to create two go-routine
	var wg sync.WaitGroup
	wg.Add(2)
	defer func() { wg.Wait() }()

	var locked int32

	go func() {
		// use CompareAndSwapInt32 to acquire a lock
		atomic.CompareAndSwapInt32(&locked, 0, 1)
		defer func() {
			// release lock
			atomic.CompareAndSwapInt32(&locked, 1, 0)
			log.Println("lock release")
			wg.Done()
		}()

		// do something
		for i := 0; i < 10; i++ {
			time.Sleep(time.Millisecond * 200)
			log.Println("do something: ", i)
		}
	}()

	go func() {
		// let the other go-routine go first
		time.Sleep(time.Second)

		// try to get the lock without rest.
		// this will never give up the cpu and in-turn block the other go-routine causing a dead-lock.
		swaped := false
		for !swaped {
			swaped = atomic.CompareAndSwapInt32(&locked, 0, 1)

			// i think the solution might be to make the go-routine sleep for a while...
			// uncomment this line and everything will be all right.
			//time.Sleep(time.Millisecond * 10)
		}

		// finally get the lock.
		// unfortunately we'll never get here.
		log.Println("swaped done")
		wg.Done()
	}()
}
