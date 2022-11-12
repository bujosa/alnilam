package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup
var counter int
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

func incrementor(s string) {
	for i := 0; i < 20; i++ {
		time.Sleep(time.Duration(rand.Intn(20))*time.Millisecond)
		mutex.Lock() // I use this because lock block the access 
		counter++
		mutex.Unlock()
	}
	wg.Done()
}

