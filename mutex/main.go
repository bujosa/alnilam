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
var mutex sync.Mutex

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Println(runtime.NumCPU())
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
		time.Sleep(time.Duration(rand.Intn(20))*time.Millisecond)
		mutex.Lock()
		x := counter
		x++
		counter = x
		mutex.Unlock()
	}
	wg.Done()
}

