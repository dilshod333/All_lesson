package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func Sorting(slice *[]int) *[]int {
	for i := 0; i < len(*slice); i++ {
		for j := i + 1; j < len(*slice); j++ {
			if (*slice)[i] > (*slice)[j] {
				(*slice)[i], (*slice)[j] = (*slice)[j], (*slice)[i]
			}
		}
	}
	return slice
}

func main() {
	map1 := make(map[int]bool)
	slice := []int{}
	n := 0
	max, min := 50, 1
	for i := 0; true; i++ {
		n = rand.Intn(max-min+1) + 1
		if map1[n] == true {
			continue
		} else {
			map1[n] = true
			if len(slice) < 20 {
				slice = append(slice, n)
			} else if len(slice) == 20 {
				break
			}
		}

	}
	for true {
		fmt.Println("20 unique random numbers :", slice)
		fmt.Print("Do you want to sort random numbers : Y/N : ")
		a := ""
		fmt.Scanln(&a)
		a = strings.ToUpper(a)
		if a == "Y" {
			sort := Sorting(&slice)
			fmt.Println("After Sorting ", *sort)
			break

		} else {
			fmt.Println("Okey bro !!!")
			break
		}
	}

}
