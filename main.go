package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	chint := make(chan int)
	wg := sync.WaitGroup{}

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			chint <- 5 * 2
		}(&wg)
	}

	go func(wg *sync.WaitGroup) {
		wg.Wait()
		close(chint)
	}(&wg)

	go func() {
		for v := range chint {
			fmt.Println(v)
		}
	}()
	time.Sleep(1 * time.Second)
}
