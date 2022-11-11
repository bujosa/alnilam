package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

// I use this packages because i need to add more go routines
// So I need to wait for this.
var wg sync.WaitGroup
var counter int

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	wg.Add(2)
	go incrementor("First:") // Call 1
	go incrementor("Second:") // Call 2
	wg.Wait()
	fmt.Println("Result: ", counter)
}

func incrementor(s string) {
	for i := 0; i < 20; i++ {
		x := counter
		x++
		time.Sleep(time.Duration(rand.Intn(3))*time.Millisecond)
		counter = x
		fmt.Println(s, i, "Counter:", counter)
	}
}

