package main

import (
	"fmt"

)

type Names struct {
	name1  string
	name2  string
	age    int
	height float32
}

func number(s, b *int) (*int, *int) {
	fmt.Println(s, b)
	*s = 5
	*b = 0

	fmt.Println("After changing the value ", *s, *b)
	return s, b
}

func main() {

	var k Names
	fmt.Printf("Type of k = %T\n", k)
	k.name1 = "Dilshod"
	fmt.Println(k)
	a, b := 3, 4
	n, n2 := number(&a, &b)
	fmt.Printf("type of n = %T and values = %v\n", n, *n)
	// fmt.Println(a, b)
	*n = 1
	fmt.Println(*n, *n2)
	addition := func(a, b int) int {
		return a + b
	}

	num := addition(3, 4)
	num = 99
	fmt.Printf("type = %T value = %d ", num, num)
	newStr := "Salom"
	newStr = string(newStr)
	fmt.Println(newStr)
}


func insertSort(a, b int)int{
	for{
		
	}
}