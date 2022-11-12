package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

var wg sync.WaitGroup
var counter int64
var mutex sync.Mutex

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

// Example using atomicity
func incrementor(s string) {
	for i := 0; i < 20; i++ {
		time.Sleep(time.Duration(rand.Intn(20))*time.Millisecond)
		atomic.AddInt64(&counter, 1)
	}
	wg.Done()
}
