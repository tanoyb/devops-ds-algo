package main

import (
	"fmt"
)

func main() {
	a := "example of string"
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	bs := []byte(a) //convert to slice of bytes
	fmt.Println(bs)
	fmt.Printf("%T\n", bs)
	for i := 0; i < len(a); i++ {
		fmt.Printf("%#U\n", a[i])
	}
	for i, v := range a {
		fmt.Println(i, v)
	}
}
