package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i // Introduce value into the channel
		}
	}()

	go func() {
		for {
			fmt.Println(<-c) // Get Channel value
		}
	}()

	time.Sleep(time.Second)
}