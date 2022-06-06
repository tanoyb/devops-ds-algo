package main

import (
	"fmt"
	"sync"
)

//Condition Wait -  put a goroutine on hold and release lock
//Condition Singal - singals the onhold goroutine
//Condition Broadcast - signals all the hold gocoutines

var sharedRsc = make(map[string]interface{})

func main() {
	var wg sync.WaitGroup
	mu := sync.Mutex{}
	cond := sync.NewCond(&mu)

	wg.Add(1)
	go func() {
		defer wg.Done()

		//TODO: suspend goroutine until sharedRsc is populated.
		cond.L.Lock()
		for len(sharedRsc) == 0 {
			//time.Sleep(1 * time.Millisecond)
			cond.Wait() //release the lock and suspend goroutine
		}

		fmt.Println(sharedRsc["rsc1"])
		cond.L.Unlock()
	}()

	// writes changes to sharedRsc
	cond.L.Lock()
	sharedRsc["rsc1"] = "foo"
	cond.Signal()
	cond.L.Unlock()

	wg.Wait()
}
