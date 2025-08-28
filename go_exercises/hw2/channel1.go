package main

import (
	"fmt"
	"time"
)

func sendData(ch chan<- int) {
	for i := 1; i <= 10; i++ {
		ch <- i
		fmt.Printf("Sent: %d\n", i)
	}
}

func receiveData(ch <-chan int) {
	for i := range ch {
		fmt.Printf("Received: %d\n", i)
	}
}

func main() {
	ch := make(chan int)

	go sendData(ch)
	go receiveData(ch)

	time.Sleep(1 * time.Second) // allow sending and receiving to complete
	close(ch)
}
