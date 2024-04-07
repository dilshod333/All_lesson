package main

import (
	"fmt"
)

type P struct {
	name   string
	age    int
	height float32
}


func main() {
	type apple struct{

	}
	
	var person P = P{
		name:   "Dilshod",
		age:    18,
		height: 1.87,
	}

	var person2 P
	person2.name = "ali"
	person2.age = 19
	person2.height = 1.90
	fmt.Println(person)
	fmt.Println(person2)
	fmt.Printf("type = %T ", person2)
}
