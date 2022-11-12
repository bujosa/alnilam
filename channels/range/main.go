package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i // Introduce value into the channel
		}

		close(c)
	}()

	// Iterate the channel until this close
	for n := range c {
		fmt.Println(n)
	}

}