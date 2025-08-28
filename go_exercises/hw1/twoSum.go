package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	ans := []int{}
	val2ind := make(map[int]int)

	for ind, val := range nums {
		val2ind[val] = ind
	}

	for ind, val := range nums {
		if _, ok := val2ind[target-val]; ok && val2ind[target-val] != ind {
			ans = append(ans, ind)
			ans = append(ans, val2ind[target-val])
			break
		}
	}

	return ans
}

func main() {
	// test 1
	nums1 := []int{2, 7, 11, 15}
	target1 := 9
	inds1 := twoSum(nums1, target1)
	fmt.Printf("nums: %v, target = %d, indices: %v\n", nums1, target1, inds1)

	// test 2
	nums2 := []int{3, 2, 4}
	target2 := 6
	inds2 := twoSum(nums2, target2)
	fmt.Printf("nums: %v, target = %d, indices: %v\n", nums2, target2, inds2)

	// test 3
	nums3 := []int{3, 3}
	target3 := 6
	inds3 := twoSum(nums3, target3)
	fmt.Printf("nums: %v, target = %d, indices: %v\n", nums3, target3, inds3)
}
