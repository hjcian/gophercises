package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"
)

func consumer(wid int, q <-chan int) {
	for jid := range q {
		r := rand.Intn(10)
		fmt.Printf("[wid:%v] [job-%v] start \n", wid, jid)
		time.Sleep(time.Duration(r) * time.Second)
		fmt.Printf("[wid:%v] [job-%v] done (spent: %v s) \n", wid, jid, r)
	}
	fmt.Printf("[wid:%v] [closed] \n", wid)
}

func concurrentProduce(totalJob int, queue chan<- int) {
	for i := 0; i < totalJob; i++ {
		go func(jid int, q chan<- int) {
			q <- jid
		}(i, queue)
	}
}

func multipleConsume(queue <-chan int) {
	n := runtime.GOMAXPROCS(-1) // get the number of CPU cores
	fmt.Printf("Total consumers: %v \n", n)
	for i := 0; i < n; i++ {
		go consumer(i, queue)
	}
}

func chanMain() {
	totalJob := 10
	queue := make(chan int, totalJob)
	multipleConsume(queue)

	go concurrentProduce(totalJob, queue)
	sigsHandler(queue)
	for {
	}
}

func sigsHandler(queue chan int) {
	intCnt := 0
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	for {
		select {
		case s := <-sigs:
			if intCnt == 0 {
				fmt.Println(">>> Got signal: ", s)
				fmt.Println(">>> close the channel")
				close(queue)
				intCnt++
			} else {
				os.Exit(0)
			}
		}
	}
}

func main() {
	chanMain()
}
