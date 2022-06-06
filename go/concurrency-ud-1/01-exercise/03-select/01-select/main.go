package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	ch3 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "one"
	}()

	go func() {
		time.Sleep(3 * time.Second)
		ch2 <- "two"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		ch3 <- "three"
	}()

	// TODO: multiplex recv on channel - ch1, ch2
	for i := 0; i < 3; i++ { //because we need to get 3 values from 3 channels
		select {
		case m := <-ch1:
			fmt.Println(m)
		case m2 := <-ch2:
			fmt.Println(m2)
		case m3 := <-ch3:
			fmt.Println(m3)
		}

	}

}
