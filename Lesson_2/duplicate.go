package main

import "fmt"

func mergeSort(s [4]int, s2 [5]int) {

	a := append(s[:], s2[:]...)

	fmt.Println(a)
}

func main() {
	db := []int{1, 2, 3, 4, 5, 2, 3}

	mp := make(map[int]int)
	for _, val := range db {
		mp[val]++
	}
	for key, val := range mp {
		if val >= 2 {
			fmt.Println("duplicate: ", key)
		} else if val == 1 {
			fmt.Printf("not duplicate %d\n", key)
		}
	}

	mergeSort([4]int{3, 2, 1, 5}, [5]int{56, 34, 23, 1, 11})
}
