package main

import (
	"fmt"
)

func singleNumber(nums []int) int {
	count := make(map[int]int)
	for _, value := range nums {
		if count[value] == 0 {
			count[value] += 1
		} else {
			count[value] = 0
		}
	}

	v := -1
	for num, cnt := range count {
		if cnt == 1 {
			v = num
			break
		}
	}

	return v
}

func main() {
	// test
	arr1 := []int{0, 0, 1, 2, 1, 10, 2}
	v1 := singleNumber(arr1)
	fmt.Printf("The signal number in %v is: %d\n", arr1, v1)
}
