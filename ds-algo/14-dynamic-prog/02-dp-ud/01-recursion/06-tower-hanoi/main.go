package main

import "fmt"

func TowerOfHanoi(n int, FROM, HELPER, TO string) {

	//take n-1 disks from tower FROM to HELPER via TO
	//then move the nth disk from FROM to TO
	//then move n-1 disks from HELPER to TO via FROM
	if n == 0 {
		return
	}
	TowerOfHanoi(n-1, FROM, TO, HELPER)    //take n-1 disks from tower A to B via C
	fmt.Println("from ", FROM, " to ", TO) // move nth disk from tower A to C
	TowerOfHanoi(n-1, HELPER, FROM, TO)    //move n-1 disks from tower B to C

}

func main() {
	TowerOfHanoi(5, "A", "B", "C")
}
