package main

import (
	"fmt"
	"time"
)

func simpleMain() {
	c := make(chan int, 10)
	go func() {
		fmt.Println("start loop")
		for v := range c {
			fmt.Println("got", v)
		}
	}()
	c <- 123
	time.Sleep(1 * time.Second)
	c <- 456
	time.Sleep(1 * time.Second)
	c <- 789
	time.Sleep(1 * time.Second)
	fmt.Println("exit")
}

func isOkMain() {
	c := make(chan int, 10)
	go func() {
		fmt.Println("start loop")
		for {
			v, ok := <-c
			if !ok {
				fmt.Println("c has been closed")
				fmt.Println("but you can still got the v:", v)
				break
			}
			fmt.Println("got", v)
		}
	}()
	c <- 123
	c <- 456
	close(c)
	// c <- 789 // send a value to closed channell will result in panic
	time.Sleep(1 * time.Second)
	fmt.Println("exit")
}

func main() {
	simpleMain()
	fmt.Println("===========================================")
	isOkMain()
}
