package main

import (
	"fmt"
	"time"
)

func receiveOnly(ch <-chan int) {
	for v := range ch {
		fmt.Printf("received: %d\n", v)
		time.Sleep(5 * time.Millisecond) // simulate processing time so that the buffer can be filled up
	}
}

// 只发送channel的函数
func sendOnly(ch chan<- int) {
	for i := 0; i < 100; i++ {
		ch <- i
		fmt.Printf("sent: %d\n", i)
	}
	close(ch)
}

func main() {
	ch := make(chan int, 5) // buffer = 5
	go sendOnly(ch)
	go receiveOnly(ch)
	time.Sleep(1 * time.Second) // wait for goroutines to finish
}
