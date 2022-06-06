package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	numOfCPU := runtime.NumCPU()
	if numOfCPU > 4 {
		numOfCPU = 4
	}
	runtime.GOMAXPROCS(numOfCPU)

	var counter uint64
	var wg sync.WaitGroup

	// TODO: implement concurrency safe counter

	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for c := 0; c < 1000; c++ {
				atomic.AddUint64(&counter, 1)
				//counter++
			}
		}()
	}
	wg.Wait()
	fmt.Println("counter: ", counter)
}

//go run -race  main.go
