package main

import "fmt"

func main() {
	var map1 = make(map[string]int)
	map1["hello"] = 12
	fmt.Println(map1["hello"])
}
