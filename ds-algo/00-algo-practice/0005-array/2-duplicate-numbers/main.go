package main

import "fmt"

func main() {
	// take a hash map
	arr := []int{1, 2, 4, 5, 6, 7, 8, 3, 1, 4, 9, 10, 9}
	hash := make(map[int]int)
	for _, v := range arr {
		//fmt.Println(v)
		if _, ok := hash[v]; ok {
			//do something here
			fmt.Println("duplicate found ", v)

		} else {
			hash[v] = 1
		}
	}

}
