package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int, 10)
	go func() {
		for {
			select {
			case <-time.After(time.Second):
				fmt.Println("put data in ch1")
				ch1 <- 123
			case v := <-ch1:
				fmt.Println("ch1:", v)
			}
		}
	}()
	select {}
}
