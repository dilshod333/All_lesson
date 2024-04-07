package main

import "fmt"

func main() {

	// anonymous function ishlatdim va two pointer technique
	s := ""
	fmt.Print("Enter the word: ")
	fmt.Scanln(&s)
	isPalindrome := func(s string) bool {
		l := 0
		r := len(s) - 1
		for l < r {
			if s[l] == s[r] {
				l++
				r--
			} else {
				return false
			}
		}
		return true
	}
	fmt.Printf("result: %v ", isPalindrome(s))

}
