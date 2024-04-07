package main

import "fmt"

func main() {
	fmt.Println(len("Hello, World"))
	fmt.Println("Hello, World"[1])
	fmt.Println("Hello, " + "World")
	a := "Salom"
	fmt.Println(string(a[0]))

	var b string = "Dilshod"
	fmt.Println(string(b[1]), len(b))

	var x string
	x = "Salom "
	x += `Hello`
	fmt.Println(x)
}
