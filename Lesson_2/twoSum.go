package main

import "fmt"

func twoSum(slice []int, target int) (int, int, int, int) {
	for i := 0; i < len(slice); i++ {
		for j := i + 1; j < len(slice); j++ {
			if slice[i]+slice[j] == target {
				return slice[i], slice[j], i, j
			}
		}
	}
	return -1, -1, -1, -1
}
func main() {
	target := 0
	fmt.Print("Enter the target number: ")
	fmt.Scanln(&target)

	var a int
	fmt.Printf("How many number you want to input: ")
	fmt.Scanln(&a)
	slice := make([]int, 0)
	for i := 1; i <= a; i++ {
		fmt.Printf("Enter the %v - number ", i)
		num := 0
		fmt.Scanln(&num)
		slice = append(slice, num)
	}

	res1, res2, idx, idj := twoSum(slice, target)
	fmt.Println("-------------------------------------")
	fmt.Printf("The two sum answer is (%v and  %v) and indexes are (%v and %v)", res1, res2, idx, idj)
}
