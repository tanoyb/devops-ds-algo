package main

import "fmt"

//if a recursive function calling iteself more than one time
//then it is called tree recursion

func fun(n int) {
	if n > 0 {
		fmt.Printf("%d\n", n)
		fun(n - 1)
		fun(n - 1)
		//trr recursion- more than one recursive call
	}
}

func main() {
	fun(3)

}
