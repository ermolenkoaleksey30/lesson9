package main

import (
	"fmt"
	"time"
)

func main() {
	chInt1 := make(chan int)
	chInt2 := make(chan int)

	go func() {
		var i int
		for {
			chInt1 <- i
			i++
			time.Sleep(1 * time.Second)
		}
	}()
	go func() {
		var i int
		for {
			chInt2 <- i
			i++
			time.Sleep(1 * time.Second)
		}
	}()

	go func() {
		for {
			select {
			case val1 := <-chInt1:
				fmt.Println("GO1", val1)
			case val2 := <-chInt2:
				fmt.Println("GO1", val2)
			case <-time.After(2 * time.Second):
				return
			}
		}
	}()

	time.Sleep(5 * time.Second)
}
