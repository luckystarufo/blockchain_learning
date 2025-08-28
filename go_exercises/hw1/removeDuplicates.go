package main

import (
	"fmt"
)

func removeDuplicates(nums []int) int {
	i, j := 0, 0
	n := len(nums)
	if n == 0 {
		return 0
	}

	for {
		v := nums[j]
		j++
		for {
			if j < n && nums[j] == v {
				j++
			} else {
				break
			}
		}
		nums[i] = v
		i++

		if j >= n {
			break
		}
	}

	return i
}

func main() {
	// test 1
	nums1 := []int{1, 1, 2}
	fmt.Printf("Before %v has length %d\n", nums1, len(nums1))
	l1 := removeDuplicates(nums1)
	fmt.Printf("After %v has length %d\n", nums1, l1)

	// test 2
	nums2 := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Printf("Before: %v has length %d\n", nums2, len(nums2))
	l2 := removeDuplicates(nums2)
	fmt.Printf("Before %v has length %d\n", nums2, l2)
}
