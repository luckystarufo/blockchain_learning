package main

import (
	"fmt"
)

func plusOne(digits []int) []int {
	ans := []int{}

	// digit by digit
	carry := 1
	for i := len(digits) - 1; i >= 0; i-- {
		cur := digits[i] + carry
		if cur == 10 {
			carry = 1
			ans = append(ans, 0)
		} else {
			carry = 0
			ans = append(ans, cur)
		}
	}

	// last digit
	if carry == 1 {
		ans = append(ans, 1)
	}

	// reverse
	n := len(ans)
	for i := 0; i < n/2; i++ {
		tmp := ans[i]
		ans[i] = ans[n-1-i]
		ans[n-1-i] = tmp
	}

	return ans
}

func main() {
	// test 0
	digits0 := []int{1, 2, 3}
	ans0 := plusOne(digits0)
	fmt.Printf("%v add one is %v\n", digits0, ans0)

	// test 1
	digits1 := []int{9}
	ans1 := plusOne(digits1)
	fmt.Printf("%v add one is %v\n", digits1, ans1)

	// test 3
	digits2 := []int{9, 9}
	ans2 := plusOne(digits2)
	fmt.Printf("%v add one is %v\n", digits2, ans2)

	// test 4
	digits3 := []int{1, 0, 9}
	ans3 := plusOne(digits3)
	fmt.Printf("%v add one is %v\n", digits3, ans3)
}
