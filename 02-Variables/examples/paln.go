package main

import (
	"fmt"
	"strings"
)

// helper function to prepare the string for palindrome check
func prepareString(s string) string {
	s = strings.ToLower(s)
	s = strings.ReplaceAll(s, " ", "")
	s = strings.ReplaceAll(s, ",", "")
	s = strings.ReplaceAll(s, "!", "")
	return s
}

// isPalindrome checks if a string is a palindrome.
func isPalindrome(s string) bool {
	s = prepareString(s)
	length := len(s)
	for i := 0; i < length/2; i++ {
		if s[i] != s[length-i-1] {
			return false
		}
	}
	return true
}

func main() {
	input := "A man, a plan, a canal, Panama!"
	input1 := "racecar"
	input2 := "hello"

	fmt.Println(isPalindrome(input))  // True
	fmt.Println(isPalindrome(input1)) // True
	fmt.Println(isPalindrome(input2)) // False
}
