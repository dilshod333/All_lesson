package main

import (
	"fmt"
)

// swaping to elements
type Person struct {
	name   string
	age    int
	height float32
}

func main() {
	// working with struct and pointer together

	var person Person = Person{
		name:   "hello",
		age:    19,
		height: 1.89,
	}
	fmt.Println(person)
	s := []int{3, 456, 6, 56}
	fmt.Println(s[:])

	var s2 = []int{3, 4, 5, 6}
	fmt.Println(s2[3:])

	number := 5
	number2 := 9
	fmt.Println("before changing: ", number, number2)
	a, a2 := swapElement(&number, &number2)
	fmt.Printf("afer changing (%v %v)\n", *a, *a2)
	k, k2 := 7, 9
	swapWithout(&k, &k2)
}

func swapElement(a, a2 *int) (*int, *int) {
	temp := *a
	*a = *a2
	*a2 = temp
	return a, a2
}

func swapWithout(n, n2 *int) {
	fmt.Println("Before changing the elements: ", *n, *n2)
	*n = *n + *n2  // n = 9
	*n2 = *n - *n2 // n2 = 4
	*n = *n - *n2

	fmt.Print("After changing elements: ", *n, *n2, "\n")

}
