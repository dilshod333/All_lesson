package main

import "fmt"

func main() {
	count := 0
	nums := []int{21, 34, 56, 54, 3456, 432}
	for _, num := range nums {
		if check(num)%2 == 0 {
			count++
		}
	}
	fmt.Println(count)

}

func check(num int) int {
	count := 0

	for num != 0 {
		num = num / 10
		count++
	}

	return count
}
