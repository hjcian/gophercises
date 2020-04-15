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

func v2(a []int, b []int, v int) bool {
	type empty struct{}
	var keyExists = empty{}

	set := make(map[int]empty)
	for _, vb := range b {
		set[vb] = keyExists
	}
	for _, va := range a {
		target := v - va
		if _, ok := set[target]; ok {
			return true
		}
	}
	return false
}

func v3(a []int, b []int, v int, n int) bool {
	type empty struct{}
	var keyExists = empty{}

	set := make(map[int]empty)
	for _, vb := range b {
		set[vb] = keyExists
	}

	setA := make(map[int]empty)
	for _, v := range a {
		setA[v] = keyExists
	}

	bucket := int(math.Ceil(float64(len(a)) / float64(n)))
	ans := false
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
			for _, va := range suba {
				target := v - va
				if _, ok := set[target]; ok {
					ans = true
				}
			}
		}()
	}
	wg.Wait()
	return ans
}
