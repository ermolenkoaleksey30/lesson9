package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var num1 int
var num2 atomic.Int64
var mtx sync.Mutex

func main() {
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		mtx.Lock()
		num1 += 2
		mtx.Unlock()
	}(wg)

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		num2.Add(1)
	}(wg)

	wg.Wait()
	fmt.Println(num1, num2.Load())
}
