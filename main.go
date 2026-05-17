package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	paretnContext, paretnContextCancle := context.WithTimeout(context.Background(), 5*time.Second)
	defer paretnContextCancle()

	go func(ctx context.Context) {
		var i int
		for {
			select {
			case <-ctx.Done():
				return
			default:
				time.Sleep(200 * time.Millisecond)
				fmt.Println(i)
				i++
			}
		}
	}(paretnContext)

	time.Sleep(8 * time.Second)
}
