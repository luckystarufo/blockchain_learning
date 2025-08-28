package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(x int) bool {
	s := strconv.Itoa(x)
	n := len(s)
	for i := 0; i < n/2; i++ {
		if s[i] != s[n-1-i] {
			return false
		}
	}

	return true
}

func main() {
	// test 0
	ans0 := isPalindrome(0)
	fmt.Printf("Is 0 palindrome? %t\n", ans0)
	// test 1
	ans1 := isPalindrome(121)
	fmt.Printf("Is 121 palindrome? %t\n", ans1)
	// test 2
	ans2 := isPalindrome(10)
	fmt.Printf("Is 10 palindrome? %t\n", ans2)
	// test 3
	ans3 := isPalindrome(1111211)
	fmt.Printf("Is 1111211 palindrome? %t\n", ans3)
	// test 3
	ans4 := isPalindrome(1112111)
	fmt.Printf("Is 1112111 palindrome? %t\n", ans4)
}
