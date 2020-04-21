package main

import (
	"log"
	"math/rand"
	"os"
	"os/signal"
	"runtime"
	"sync"
	"syscall"
	"time"
)

func consumer(wid int, q <-chan int, wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()
	for jid := range q {
		r := rand.Intn(3)
		log.Printf("[wid:%v] [job-%v] start \n", wid, jid)
		time.Sleep(time.Duration(r) * time.Second)
		log.Printf("[wid:%v] [job-%v] done (spent: %v s) \n", wid, jid, r)
	}
	log.Printf("[wid:%v] [closed] \n", wid)
}

func concurrentProduce(totalJob int, queue chan<- int) {
	for i := 0; i < totalJob; i++ {
		go func(jid int, q chan<- int) {
			q <- jid
		}(i, queue)
	}
	log.Println("producer done")
}

func multipleConsume(queue <-chan int, wg *sync.WaitGroup) {
	n := runtime.GOMAXPROCS(-1) // get the number of CPU cores
	log.Printf("Total consumers: %v \n", n)
	for i := 0; i < n; i++ {
		go consumer(i, queue, wg)
	}
}

func chanMain() {
	totalJob := 10
	queue := make(chan int, totalJob)
	wg := sync.WaitGroup{}
	multipleConsume(queue, &wg)

	go concurrentProduce(totalJob, queue)
	sigsHandler(queue, &wg)
}

func sigsHandler(queue chan int, wg *sync.WaitGroup) {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM) // this method will cause a infinite loop
	for s := range sigs {
		log.Println(">>> Got signal: ", s)
		log.Println(">>> close the channel")
		close(queue)
		wg.Wait()
		log.Println(">>> close the sigs")
		close(sigs)
	}
}

func main() {
	chanMain()
}
