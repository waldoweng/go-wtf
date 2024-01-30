package main

import (
	"flag"
	"fmt"
	"runtime"
	"sync"
	"time"

	"golang.org/x/sys/unix"
)

var (
	nPrioritizedGoroutines = flag.Int("p", 1, "# of prioritized goroutines")
	nNormalGoroutines      = flag.Int("n", 1, "# of normal goroutines")
	restDuration           = flag.Duration("r", 0, "rest for a certain amount of time between works")
)

func prioritizeThread() {
	// set thread priority to the highest
	a := unix.SchedAttr{
		Size:     unix.SizeofSchedAttr,
		Policy:   1,
		Priority: 49,
	}
	if err := unix.SchedSetAttr(0, &a, 0); err != nil {
		panic(err)
	}
}

func doWorks(workerId int) {
	t := time.Now()
	for i := 0; i < 100; i++ {
		st := time.Now()
		res := 0
		for ii := 0; ii < 1e9; ii++ {
			res += ii
		}
		fmt.Printf("%d@%d, timecost: %s, res: %d \n", workerId, unix.Gettid(), time.Since(st), res)

		// sleep for a while to simulate gaps between requests.
		if *restDuration > 0 {
			time.Sleep(*restDuration)
		}
	}
	fmt.Printf("total execute time for worker: %d is %s\n", workerId, time.Since(t))
}

func main() {
	flag.Parse()

	runtime.GOMAXPROCS(*nPrioritizedGoroutines + *nNormalGoroutines)
	var wg sync.WaitGroup

	workerId := 0
	for i := 0; i < *nPrioritizedGoroutines; i++ {
		wg.Add(1)
		go func() {
			// assign goroutine to a designated thread
			runtime.LockOSThread()
			// prioritize this thread
			prioritizeThread()

			defer wg.Done()
			doWorks(workerId)
		}()
		workerId++
	}

	for i := 0; i < *nNormalGoroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			doWorks(workerId)
		}()
		workerId++
	}

	wg.Wait()
}

