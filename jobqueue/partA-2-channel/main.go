package main

import (
	"fmt"
	"time"
)

func producer(jid int, q chan<- int) {
	q <- jid
	fmt.Println("Add job", jid)
}

func consumer(q <-chan int) {
	jid := <-q
	fmt.Println("Do job", jid)
}

func concurrentProduce(totalJob int, queue chan<- int) {
	for i := 0; i < totalJob; i++ {
		go producer(i, queue)
	}
}

func consume(queue <-chan int) {
	for {
		consumer(queue)
	}
}

func chanMain() {
	totalJob := 10
	queue := make(chan int, totalJob)
	fmt.Println("initial queue len: ", len(queue), "cap:", cap(queue))

	go concurrentProduce(totalJob, queue)
	go consume(queue)

	fmt.Println("main goroutine go to sleep...")
	time.Sleep(time.Second * 3)

	fmt.Println("final queue len:", len(queue), "cap:", cap(queue))
}

func main() {
	chanMain()
}
