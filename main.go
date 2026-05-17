package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var num1 atomic.Int64
var num2 int = 0
var mtx sync.Mutex

func main() {
	wg := &sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			num1.Add(1)
		}()
	}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mtx.Lock()
			num2 = num2 + 2
			mtx.Unlock()
		}()
	}

	wg.Wait()
	fmt.Println(num1.Load(), num2)
}
