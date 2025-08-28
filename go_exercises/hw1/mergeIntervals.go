package main

import (
	"fmt"
	"sort"
)

func merge(intervals [][]int) [][]int {
	n := len(intervals)
	if n == 0 {
		return intervals
	}

	// sort according to the start element first, then end element
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] < intervals[j][1]
		}
		return intervals[i][0] < intervals[j][0]
	})

	ans := [][]int{}
	cur := []int{intervals[0][0], intervals[0][1]}
	for k := 1; k < n; k++ {
		if intervals[k][0] <= cur[1] {
			// merge the overlapped intervals
			cur = []int{cur[0], intervals[k][1]}
		} else {
			ans = append(ans, cur)
			cur = []int{intervals[k][0], intervals[k][1]}
		}
	}

	ans = append(ans, cur)

	return ans
}

func main() {
	// test 1
	intervals1 := [][]int{{2, 6}, {1, 3}, {8, 10}, {15, 18}}
	fmt.Printf("Before %v\n", intervals1)
	intervals_merged1 := merge(intervals1)
	fmt.Printf("After %v\n", intervals_merged1)

	// test 2
	intervals2 := [][]int{{1, 4}, {4, 5}}
	fmt.Printf("Before %v\n", intervals2)
	intervals_merged2 := merge(intervals2)
	fmt.Printf("After %v\n", intervals_merged2)
}
