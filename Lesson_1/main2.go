package main

import (
	"fmt"
)

func main() {
	a := "Hey brooo h hoow  hh"
	fmt.Printf("Salom    heyyyyy       bro ", "Hello ", "Hi ", 10, 23.45, true, "\n")
	fmt.Print("Salom \nhi\neee\nlo ", a)
	_ = 77
	x := 19
	y := &x
	*y = 01
	fmt.Println(x, y)
}
