package sumoftwo

import (
	"math"
	"sync"
)

func v1(a []int, b []int, v int) bool {
	for i := range a {
		for j := range b {
			if (a[i] + b[j]) == v {
				return true
			}
		}
	}
	return false
}

type empty struct{}

var keyExists = empty{}

func search(given []int, set map[int]empty) bool {
	for _, va := range given {
		if _, ok := set[va]; ok {
			return true
		}
	}
	return false
}

func v2(a []int, b []int, v int) bool {
	set := make(map[int]empty)
	for _, vb := range b {
		set[v-vb] = keyExists
	}
	return search(a, set)
}

// concurrent version: wait for all goroutine finished
func v3(a []int, b []int, v int, n int) bool {
	set := make(map[int]empty)
	for _, vb := range b {
		set[v-vb] = keyExists
	}

	bucket := int(math.Ceil(float64(len(a)) / float64(n)))

	ansCh := make(chan empty, 1)
	wg := sync.WaitGroup{}
	wg.Add(n)
	for i := 0; i < n; i++ {
		high := (i + 1) * bucket
		if high > len(a) {
			high = len(a)
		}
		suba := a[i*bucket : high]
		go func() {
			defer wg.Done()
			isHit := search(suba, set)
			if isHit && len(ansCh) == 0 {
				ansCh <- keyExists
			}
		}()
	}
	wg.Wait()
	return len(ansCh) == 1
}
