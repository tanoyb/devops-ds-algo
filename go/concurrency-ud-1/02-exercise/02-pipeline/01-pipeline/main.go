package main

import (
	"fmt"
)

// TODO: Build a Pipeline
// generator() -> square() -> print
// 1st stage    2nd stage	  3rd stage
// generator - convertes a list of integers to a channel
func generator(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, v := range nums {
			out <- v
		}
		close(out)
	}()

	return out
}

// square - receive on inbound channel
// square the number
// output on outbound channel
func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for v := range in {
			out <- v * v
		}
		close(out)
	}()
	return out
}

func main() {
	// set up the pipeline

	// run the last stage of pipeline
	// receive the values from square stage
	// print each one, until channel is closed.
	ch := generator(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	out := square(ch)

	for n := range out {
		fmt.Println(n)
	}
}
