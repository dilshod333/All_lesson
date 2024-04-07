package main

import (
	"fmt"
	"math/rand"
)

func main() {
	randomInt := rand.Intn(51)
	fmt.Println("random int ", randomInt)
	if randomInt > 1 && randomInt < 10 {
		fmt.Println("its number between 1 to 10:", randomInt)
	} else if randomInt > 10 && randomInt < 20 {
		fmt.Println("its number between 10 to 20:", randomInt)
	} else if randomInt > 20 && randomInt < 30 {
		fmt.Println("its number between 20 to 30:", randomInt)
	} else if randomInt > 30 && randomInt < 40 {
		fmt.Println("its number between 30 to 40:", randomInt)
	} else {
		fmt.Println("its number between 40 to 50:", randomInt)
	}
}
