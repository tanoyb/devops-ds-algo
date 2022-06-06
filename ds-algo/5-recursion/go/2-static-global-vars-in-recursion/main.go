package main

import "fmt"

//there is no static variable in golang like C/C++
//we will declare a package level variable to maintain a
//single copy of the variable in RAM, so that it simulates
//the behavior of a static variable
//x is actually a global variable here
var x int = 0 //one copy of x throughout the program

func fun(n int) int { //returns sum of n numbers
	if n > 0 {
		x++
		fmt.Println("x == ", x)
		return fun(n-1) + x //head recursion, adds x in return time
		//trace the function by each step to uderstand
		//in each recursive call, the x points to the
		//same variable, it does not create different
		//occurences of x in stach for each resursive
		//calls
	}
	return 0
}

func main() {
	s := fun(5)
	fmt.Println("sum == ", s)
	s = fun(5)
	fmt.Println("sum == ", s)

}
