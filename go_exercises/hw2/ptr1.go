package main

import (
	"fmt"
)

func addTen(a *int) {
	*a += 10
}

func main() {
	// test
	x := 5
	fmt.Println("before adding 10: ", x) // Should print 5
	addTen(&x)
	fmt.Println("after adding 10:", x) // Should print 15
}
