package main

import (
	"fmt"
	"time"
)

func main() {
	chInt1 := make(chan int)
	chInt2 := make(chan int)
	chInt3 := make(chan int)
	sliceChInt := []int{}

	go func() {
		for {
			time.Sleep(500 * time.Millisecond)
			chInt1 <- 1
		}
	}()
	go func() {
		for {
			time.Sleep(2 * time.Second)
			chInt2 <- 2
		}
	}()
	go func() {
		for {
			time.Sleep(900 * time.Millisecond)
			chInt3 <- 3
		}
	}()

	go func() {

		for {
			select {
			case val1 := <-chInt1:
				sliceChInt = append(sliceChInt, val1)
			case val2 := <-chInt2:
				sliceChInt = append(sliceChInt, val2)
			case val3 := <-chInt3:
				sliceChInt = append(sliceChInt, val3)
			}
		}
	}()

	time.Sleep(5 * time.Second)
	fmt.Println(sliceChInt)
}
