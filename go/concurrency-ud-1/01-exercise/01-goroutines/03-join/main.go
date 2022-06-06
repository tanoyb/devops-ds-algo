package main

import (
	"fmt"
	"sync"
)

func main() {

	var data int

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		data++
		wg.Done()
	}()

	wg.Wait()
	fmt.Printf("the value of data is %v\n", data)

	fmt.Println("Done..")
}
