package main

import (
	"fmt"
)

func multiplyByTwo(a *[]int) {
	for i := range *a {
		(*a)[i] *= 2
	}
}

func main() {
	// test
	arr := []int{1, 2, 3, 4, 5}
	fmt.Printf("before multiplying by 2: %v\n", arr)
	multiplyByTwo(&arr)
	fmt.Printf("after multiplying by 2: %v\n", arr)
}
