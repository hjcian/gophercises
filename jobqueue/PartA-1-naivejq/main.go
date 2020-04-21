package main

import (
	"log"
	"sync"
	"time"
)

func producer(jid int, q *[]int) {
	(*q) = append((*q), jid)
	log.Println("Add job", jid)
}

func consumer(q *[]int) {
	if len(*q) > 0 {
		jid := (*q)[0]
		(*q) = (*q)[1:len(*q)]
		log.Println("Do job", jid)
	}
}

func concurrentProduce(totalJob int, q *[]int) {
	for i := 0; i < totalJob; i++ {
		go producer(i, q)
	}
}

func consume(q *[]int) {
	for {
		consumer(q)
	}
}

func safeProduce(totalJob int, queue *[]int) {
	mu := sync.Mutex{}
	for i := 0; i < totalJob; i++ {
		go func(jid int, q *[]int) {
			mu.Lock()
			(*q) = append((*q), jid) // critical section
			mu.Unlock()
		}(i, queue)
	}
}

func naiveMain() {
	totalJob := 10
	queue := make([]int, 0, totalJob)
	log.Println("initial queue length: ", len(queue))

	go concurrentProduce(totalJob, &queue)
	go consume(&queue)

	log.Println("main goroutine go to sleep...")
	time.Sleep(time.Second * 3)

	log.Println("final queue length:", len(queue))
}

// analysisProduce observe the race condition take place in job queue (data lost)
func analysisProduce() {
	totalJob := 10
	queue := make([]int, 0, totalJob)
	log.Println(
		"initial queue length: ",
		len(queue))

	go concurrentProduce(totalJob, &queue)

	log.Println("main goroutine go to sleep...")
	time.Sleep(time.Second * 3)

	log.Println("final queue length:", len(queue))
}

// mutexMain use mutex avoid the race condition, and save the world
func mutexMain() {
	totalJob := 100000
	queue := make([]int, 0, totalJob)
	log.Println("initial queue length: ", len(queue))

	go safeProduce(totalJob, &queue)

	log.Println("main goroutine go to sleep...")
	time.Sleep(time.Second * 3)

	log.Println("final queue length:", len(queue))
}

func main() {
	naiveMain()
	// analysisProduce()
	// mutexMain()
}
