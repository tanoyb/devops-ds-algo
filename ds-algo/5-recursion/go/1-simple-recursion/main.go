package main

import "fmt"

//simple example of recursion
func fun(n int) { //tail recursion
	if n > 0 {
		fmt.Println("n == ", n) //prints first, then calls recursive function fun
		fun(n - 1)
		//in case of tail recursion loop is efficient, because
		// tail recrsion consumes more space in RAM
	}
}

//TAIL recursion : when the recursive call statement is the last statement of the function
//after the call the function is performing nothing is called tail recursion

func fun1(n int) { //head recursion
	if n > 0 {
		fun1(n - 1) //calls recursive functin fun1 first, then prints value in return cycle
		fmt.Println("n1 == ", n)
		//converting head recursion to loop is difficult
	}
}

//HEAD recursion : when the recursive call is the first statement of the function
//after the head recursive call, the function will perform more statements is a
//property of head recursion
func main() {
	fun(5)
	fun1(5)
}

//O(n)
