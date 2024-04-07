package main

import "fmt"

type mevallar struct {
	apple   string
	banana  string
	kiwi    string
	Isthere bool
}

func main() {
	var meva mevallar
	meva.apple = "olma"
	var ptr *int // ptr is a pointer to int, uninitialized
	fmt.Println(&ptr)

	var slice []int
	if slice == nil {
		fmt.Println("empty slice is equal to nil")
	}

	slice2 := []int{}
	if slice2 == nil {
		fmt.Println("This is empty slice2 also nil ")
	}

	fmt.Printf("Slice one = %v  Slice2= %v\n", slice, slice2)

}
