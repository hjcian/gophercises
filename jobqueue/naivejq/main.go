package main

import (
	"fmt"
	"sync"
	"time"
)

func producer(jid int, q *[]int) {
	(*q) = append((*q), jid)
	fmt.Println("Add job", jid)
}

func consumer(q *[]int) {
	if len(*q) > 0 {
		jid := (*q)[0]
		(*q) = (*q)[1:len(*q)]
		fmt.Println("Do job", jid)
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
	totalJob := 1000
	queue := make([]int, 0, totalJob)
	fmt.Println("initial queue length: ", len(queue))

	go concurrentProduce(totalJob, &queue)
	go consume(&queue)

	fmt.Println("main goroutine go to sleep...")
	time.Sleep(time.Second * 3)

	fmt.Println("final queue length:", len(queue))
}

// analysisProduce observe the race condition take place in job queue (data lost)
func analysisProduce() {
	totalJob := 10
	queue := make([]int, 0, totalJob)
	fmt.Println(
		"initial queue length: ",
		len(queue))

	go concurrentProduce(totalJob, &queue)

	fmt.Println("main goroutine go to sleep...")
	time.Sleep(time.Second * 3)

	fmt.Println("final queue length:", len(queue))
}

// mutexMain use mutex avoid the race condition, and save the world
func mutexMain() {
	totalJob := 100000
	queue := make([]int, 0, totalJob)
	fmt.Println("initial queue length: ", len(queue))

	go safeProduce(totalJob, &queue)

	fmt.Println("main goroutine go to sleep...")
	time.Sleep(time.Second * 3)

	fmt.Println("final queue length:", len(queue))
}

func main() {
	// naiveMain()
	// analysisProduce()
	mutexMain()
}
