package main

import (
	"context"
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

func consumer(ctx context.Context, wid int, q <-chan int) {
	for {
		select {
		case <-ctx.Done():
			log.Printf("[wid:%v] [closed] (reason: %s)\n", wid, ctx.Err())
			return
		case jid := <-q:
			r := rand.Intn(3)
			log.Printf("[wid:%v] [job-%v] start \n", wid, jid)
			time.Sleep(time.Duration(r) * time.Second)
			log.Printf("[wid:%v] [job-%v] done (spent: %v s) \n", wid, jid, r)
		}
	}
}

func multipleConsume(ctx context.Context, queue <-chan int, wg *sync.WaitGroup) {
	n := runtime.GOMAXPROCS(-1) // get the number of CPU cores
	log.Printf("Total consumers: %v \n", n)
	for i := 0; i < n; i++ {
		go func(i int) {
			wg.Add(1)
			defer wg.Done()
			consumer(ctx, i, queue)
		}(i)
	}
}

func chanMain() {
	totalJob := 10
	queue := make(chan int, totalJob)
	ctx, cancel := context.WithCancel(context.Background())
	wg := sync.WaitGroup{}
	multipleConsume(ctx, queue, &wg)

	go concurrentProduce(totalJob, queue)
	sigsHandler(queue, &wg, cancel)
}

func sigsHandler(queue chan int, wg *sync.WaitGroup, cancel context.CancelFunc) {
	sigs := make(chan os.Signal, 1)
	// signal.Notify will cause an infinite loop to listening signal event
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	for s := range sigs {
		log.Println(">>> Got signal: ", s, ", call cancel() to notify goroutines")
		cancel()
		wg.Wait() // wait for all job finished
		log.Println(">>> close the sigs (exit signal listening loop)")
		close(sigs)
	}
}

func main() {
	chanMain()
}
