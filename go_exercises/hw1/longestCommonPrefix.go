package main

import (
	"fmt"
)

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	prefix := ""
	i := 0
outer:
	for {
		var ch byte
		for j, str := range strs {
			if i == len(str) {
				break outer
			}
			if j == 0 {
				ch = str[i]
			} else if ch != str[i] {
				break outer
			}
		}
		prefix += string(ch)
		i++
	}

	return prefix
}

func main() {
	// test 1
	strs1 := []string{"flower", "flow", "flight"}
	ans1 := longestCommonPrefix(strs1)
	fmt.Printf("longest common prefix of %s is %s\n", strs1, ans1)

	// test 2
	strs2 := []string{}
	ans2 := longestCommonPrefix(strs2)
	fmt.Printf("longest common prefix of %s is %s\n", strs2, ans2)
}
