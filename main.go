package main

import (
	"fmt"
	"time"
)

func Job(ch chan int) {
	time.Sleep(5 * time.Second)
	ch <- 42
}

func main() {
	ch := make(chan int)

	go Job(ch)
	select {
	case val1 := <-ch:
		fmt.Println(val1)
	case <-time.After(1 * time.Second):
		return
	}

	go Job(ch)
	timer := time.NewTimer(1 * time.Second)
	select {
	case <-ch:
		timer.Stop()
	case <-timer.C:
		fmt.Println("timerC")
	}
}
