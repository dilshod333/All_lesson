package main

import (
	"fmt"
	"math/rand"
)

func main() {
	n := rand.Intn(50) // 0 < 49
	if n >= 0 && n <= 10 {
		fmt.Println("numbers between 0 to 10 ", n)
	} else if n >= 11 && n <= 20 {
		fmt.Println("numbers between 11 to 20 ", n)
	} else if n >= 21 && n <= 30 {
		fmt.Println("numbers between 21 to 30 ", n)
	} else if n >= 31 && n <= 40 {
		fmt.Println("numbers between 31 to 40 ", n)
	} else {

		fmt.Println("numbers 41 to 50 ", n)
	}

	n2 := rand.Intn(50)
	switch {
	case n2 >= 0 && n2 <= 10:
		fmt.Println("this number between 0 to 10 ", n2)
	case n2 >= 11 && n2 <= 20:
		fmt.Println("this number between 11 to 20 ", n2)
	case n2 >= 21 && n2 <= 30:
		fmt.Println("This number between 21  to 30 ", n2)
	case n2 >= 31 && n2 <= 40:
		fmt.Println("this number between 31 to 40 ", n2)
	default:
		fmt.Println("this number between 41 to 50 ", n2)
	}
}
