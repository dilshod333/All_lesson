package main 

import "fmt"

func main(){
	var a, b int 
	fmt.Printf("Enter the first number: ",)
	fmt.Scanln(&a)
	fmt.Printf("Enter the second number: ")
	fmt.Scanln(&b)

	fmt.Printf("First number = %d\nsecond number = %d\nHello world\n", a, b)

	fmt.Println("-------------------------------")
	// IKKINCHI USUL 

	n1, n2 := "", ""
	fmt.Print("Enter the first Word: ")
	fmt.Scanln(&n1)
	fmt.Print("Enter the second Word: ")
	fmt.Scanln(&n2)
	
	fmt.Printf("%v", n1 +" "+ n2 )
	
}