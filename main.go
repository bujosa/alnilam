package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// I use this packages because i need to add more go routines
// So I need to wait for this.
var wg sync.WaitGroup

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	wg.Add(2)
	go debian() // Call 1
	go cali() // Call 2
	wg.Wait()
}

// Second 
func debian(){
	for i := 0; i < 30; i++ {
		fmt.Println("debian:", i)
		time.Sleep(time.Duration(1 * time.Millisecond)) // This is for see the concurrency
	}
	wg.Done()
}

// Third
func cali(){
	for i := 0; i < 30; i++ {
		fmt.Println("cali:", i)
		time.Sleep(time.Duration(2 * time.Millisecond))
	}
	wg.Done()
}