package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var num1 atomic.Int64
var num2 int = 0
var mtx sync.Mutex

func main() {
	for i := 0; i < 5; i++ {
		go func() {
			num1.Add(1)
		}()
	}

	for i := 0; i < 5; i++ {
		go func() {
			mtx.Lock()
			num2 = num2 + 2
			mtx.Unlock()
		}()
	}

	time.Sleep(2 * time.Second)
	fmt.Println(num1.Load(), num2)
}
