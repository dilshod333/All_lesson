package main

import (
	"fmt"
	"reflect"
)

func anagram(s, t string) bool {

	map1 := make(map[rune]int)
	map2 := make(map[rune]int)

	for _, num := range s {
		num := rune(num)
		map1[num]++
	}
	for _, num := range t {
		num := rune(num)
		map2[num]++
	}
	a := reflect.DeepEqual(map1, map2)
	if a == true {
		return true
	} else {
		return false
	}

}

func main() {
	a := anagram("hello", "llohe")
	fmt.Println(a)
}
