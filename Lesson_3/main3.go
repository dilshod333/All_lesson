// // package main

// // import "fmt"

// // func main() {

// // 	s := "sdfghjbnjklhjkl"
// // 	for i := 0; i < len(s); i++ {
// // 		if i%2 == 0 && s[i]%2 == 0 {
// // // 			fmt.Println(string(s[i]))
// // // 		}
// // // 	}

// // package main

// // import (
// // 	"fmt"
// // )

// // func main() {
// // 	t := "sdfghjk 345 fnlknf 56.5gh 34 "

// // 	for i := 0; i < len(t); i++ {
// // 		if t[i] >= '0' && t[i] <= '9' {
// // 			if string(t[i]) != "" {
// // 				fmt.Print(string(t[i]))

// // 			}
// // 		}
// // 	}
// // }

// // // output 345 56.5 34

// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func main() {
// 	t := "sdfghjk 345 fnlknf 56.5.gh 34 "

// 	var str string

// 	for i := 0; i < len(t); i++ {
// 		if t[i] >= '0' && t[i] <= '9' || t[i] == '.' {
// 			str += string(t[i])
// 		} else {
// 			if str != "" {
// 				fmt.Println(strings.TrimSpace(str))
// 				str = ""
// 			}
// 		}
// 	}

//		// fmt.Printf(str)
//	}
package main

import (
	"fmt"
	//"strings"
)

func main() {
	t := "sdfghjk 345 fnlknf 56.5.gh 34 "
	//str := ""
	for i := 0; i < len(t); i++ {
		if string(t[i]) >= "0" && string(t[i]) <= "9" || string(t[i]) == "." && t[i+1] >= 47 && t[i+1] <= 58 {
			fmt.Print(string(t[i]))

		} else if string(t[i]) == " " {
			fmt.Print(" ")
		}
	}
}
