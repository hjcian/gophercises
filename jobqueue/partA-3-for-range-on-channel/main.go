package main

import (
	"log"
	"time"
)

func simpleMain() {
	c := make(chan int, 10)
	go func() {
		log.Println("start loop")
		for v := range c {
			log.Println("got", v)
		}
		log.Println("never reach this line")
	}()
	c <- 123
	time.Sleep(1 * time.Second)
	c <- 456
	time.Sleep(1 * time.Second)
	c <- 789
	time.Sleep(1 * time.Second)
	log.Println("exit")
}

func isOkMain() {
	c := make(chan int, 10)
	go func() {
		log.Println("start loop")
		for {
			v, ok := <-c
			if !ok {
				log.Println("c has been closed")
				log.Println("but you can still got the v:", v)
				break
			}
			log.Println("got", v)
		}
	}()
	c <- 123
	c <- 456
	close(c)
	// c <- 789 // send a value to closed channell will result in panic
	time.Sleep(1 * time.Second)
	log.Println("exit")
}

func main() {
	simpleMain()
	log.Println("===========================================")
	isOkMain()
}
