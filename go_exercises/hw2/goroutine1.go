package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		for i := 0; i < 10; i++ {
			if i%2 != 0 {
				fmt.Println("from routine 1: ", i)
			}
		}
	}()

	go func() {
		for i := 0; i < 10; i++ {
			if i%2 == 0 {
				fmt.Println("from routine 2: ", i)
			}
		}
	}()

	time.Sleep(1 * time.Second) // Wait for goroutines to finish
}
