package main

import "fmt"

func main() {
	c := incrementor()
	for n := range puller(c) {
		fmt.Println(n)
	}
}

// This <-chan is maybe optional because isnt affect the functionality
// but for the programmer is more clear to undestand
func incrementor() <-chan int{
	out := make(chan int)
	go func(){
		for i := 0; i < 10; i++{
			out <- i
		}
		close(out)
	}()

	return out
}

func puller(c <-chan int) <-chan int{
	out := make(chan int)
	go func(){
		var sum int
		for n := range c{
			sum += n
		}
		out <- sum
		close(out)
	}()
	return out
}