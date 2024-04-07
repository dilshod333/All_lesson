package main

import "fmt"

func bubbleSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	return arr
}

func main() {
	// bubble sort
	n := 0
	fmt.Print("How many number u want to input: ")
	fmt.Scanln(&n)
	slice := make([]int, 0)
	for i := 1; i <= n; i++ {
		num := 0
		fmt.Printf("Enter the %v - number: ", i)
		fmt.Scanln(&num)
		slice = append(slice, num)
	}
	result := bubbleSort(slice)
	fmt.Println(result)
}
