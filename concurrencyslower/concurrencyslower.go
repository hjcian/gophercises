package concurrencyslower

import (
	"runtime"
	"sync"
)

const (
	limit = 10000000000
)

func SerialSum() int {
	sum := 0
	for i := 0; i < limit; i++ {
		sum += i
	}
	return sum
}

func ConcurrentSum() int {
	n := runtime.GOMAXPROCS(0)

	sums := make([]int, n)

	wg := sync.WaitGroup{}
	for i := 0; i < n; i++ {

		wg.Add(1)
		go func(i int) {
			start := (limit / n) * i
			end := start + (limit / n)

			for j := start; j < end; j++ {
				sums[i] += j
			}

			wg.Done()
		}(i)
	}

	wg.Wait()

	sum := 0
	for _, s := range sums {
		sum += s
	}
	return sum
}

func ChannelSum() int {
	n := runtime.GOMAXPROCS(0)

	res := make(chan int)

	for i := 0; i < n; i++ {
		go func(i int, r chan<- int) {
			sum := 0
			start := (limit / n) * i
			end := start + (limit / n)
			for j := start; j < end; j++ {
				sum += j
			}
			r <- sum
		}(i, res)
	}

	sum := 0
	for i := 0; i < n; i++ {
		sum += <-res
	}
	return sum
}
