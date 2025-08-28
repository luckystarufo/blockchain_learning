package main

import (
	"fmt"
)

func isValid(s string) bool {
	pairs := map[rune]rune{')': '(', ']': '[', '}': '{'}
	stack := []rune{}
	for _, ch := range s {
		switch ch {
		case '(', '[', '{':
			stack = append(stack, ch) // push
		case ')', ']', '}':
			if len(stack) == 0 || stack[len(stack)-1] != pairs[ch] {
				return false
			}
			stack = stack[:len(stack)-1] // pop
		}
	}
	return len(stack) == 0
}

func main() {
	// test 0
	s0 := "()"
	ans0 := isValid(s0)
	fmt.Printf("Is %s valid? %t\n", s0, ans0)

	// test 1
	s1 := "()[]{}"
	ans1 := isValid(s1)
	fmt.Printf("Is %s valid? %t\n", s1, ans1)

	// test 2
	s2 := "([])"
	ans2 := isValid(s2)
	fmt.Printf("Is %s valid? %t\n", s2, ans2)

	// test 3
	s3 := "([)]"
	ans3 := isValid(s3)
	fmt.Printf("Is %s valid? %t\n", s3, ans3)

	// test 4
	s4 := "([)"
	ans4 := isValid(s4)
	fmt.Printf("Is %s valid? %t\n", s4, ans4)
}
