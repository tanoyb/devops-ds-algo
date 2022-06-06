package main

import "fmt"

//when funcA calls funcB, B calls A again.
func funA(n int) {
	if n > 0 {
		fmt.Println(n)
		funB(n - 1)
	}
}
func funB(n int) {
	if n > 1 {
		fmt.Println(n)
		funA(n / 2)
	}
}
func main() {
	funA(20)
}
