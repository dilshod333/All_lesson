package main

import "fmt"

func main() {

	map1 := make(map[int]int)
	map2 := make(map[int]int)

	slice := []int{1, 2, 3, 4, 2, 3}
	slice2 := []int{4, 7, 9, 5, 6, 87, 3}

	for _, num := range slice {
		map1[num]++
	}

	for _, num := range slice2 {
		map2[num]++
	}

	var inter []int
	for num := range map1 {

		if map2[num] > 0 {
			inter = append(inter, num)
		}
	}

	fmt.Println("slices which both have : ", inter)
}
