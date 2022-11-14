package main

import "fmt"

func main() {
	n := 10
	c := make(chan int)
	done := make(chan bool)

	for i := 0; i < n; i++ {
		go func() {
			for i := 0; i < n; i++ {
				c <- i
			}
			done <- true
		}()
	}

	// How many functions I can put in the same channels? - No number for that
	go func() {
		for i := 0; i < n; i++ {
			<-done
		}
		close(c)
	}()

	for n := range c {
		fmt.Println(n)
	}
}
