package racer

import (
	"fmt"
	"net/http"
	"time"
)

func measureResponseTime(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}

func Racer(a, b string) (winner string) {

	aDuration := measureResponseTime(a)
	bDuration := measureResponseTime(b)

	if aDuration < bDuration {
		return a
	}
	return b
}

func ConcurrentRacer(a, b string, timedOut time.Duration) (winner string, err error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timedOut):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}
}

func ping(url string) chan struct{} {
	sig := make(chan struct{})
	go func() {
		http.Get(url)
		sig <- struct{}{}
	}()
	return sig
}
