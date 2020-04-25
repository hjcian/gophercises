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

func concurrentProduce(totalJob int, queue chan<- int) {
	for i := 0; i < totalJob; i++ {
		go func(jid int, q chan<- int) {
			q <- jid
		}(i, queue)
	}
	log.Println("producer done")
}

func consumer(wid int, q <-chan int, quitSig <-chan struct{}) {
	for {
		select {
		case <-quitSig:
			log.Printf("[wid:%v] [closed] \n", wid)
			return
		case jid := <-q:
			r := rand.Intn(3)
			log.Printf("[wid:%v] [job-%v] start \n", wid, jid)
			time.Sleep(time.Duration(r) * time.Second)
			log.Printf("[wid:%v] [job-%v] done (spent: %v s) \n", wid, jid, r)
		}
	}
}

func multipleConsume(quit <-chan struct{}, queue <-chan int, wg *sync.WaitGroup) {
	n := runtime.GOMAXPROCS(-1) // get the number of CPU cores
	log.Printf("Total consumers: %v \n", n)
	for i := 0; i < n; i++ {
		go func(i int) {
			wg.Add(1)
			defer wg.Done()
			consumer(i, queue, quit)
		}(i)
	}
}

func chanMain() {
	totalJob := 10
	queue := make(chan int, totalJob)
	quitSig := make(chan struct{})
	wg := sync.WaitGroup{}
	multipleConsume(quitSig, queue, &wg)

	go concurrentProduce(totalJob, queue)
	sigsHandler(quitSig, queue, &wg)
}

func sigsHandler(quitSig chan struct{}, queue chan int, wg *sync.WaitGroup) {
	sigs := make(chan os.Signal, 1)
	// signal.Notify will cause an infinite loop to listening signal event
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	for s := range sigs {
		log.Println(">>> Got signal: ", s, ", close the quit channel avoiding panic")
		close(quitSig)
		wg.Wait() // wait for all job finished
		log.Println(">>> close the sigs (exit signal listening loop)")
		close(sigs)
	}
}

func main() {
	chanMain()
}
