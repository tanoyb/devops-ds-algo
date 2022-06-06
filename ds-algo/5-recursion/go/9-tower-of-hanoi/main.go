package main

import "fmt"

//n=number of disk, A=1st tower, B=2nd tower, C=3rd Tower
func TOH(n, A, B, C int) {
	if n > 0 {
		TOH(n-1, A, C, B)
		fmt.Printf("(Step -> Tower no %d to Tower no %d)\n", A, C)
		TOH(n-1, B, A, C)
	}
}

//we have to move a disk at a time
//no larger disk can be kept on a smaller disk
func main() {
	TOH(3, 1, 2, 3)
}
