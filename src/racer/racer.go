package racer

import (
	"net/http"
	"time"
)

func Racer(a, b string) (winner string) {
	select {
	case <-ping(a):
		return a
	case <-ping(b):
		return b
	}

}

func measureResponseTime(url string) time.Duration {
	start := time.Now()
	_, _ = http.Get(url)
	return time.Since(start)

}

func ping(url string) chan bool {
	ch := make(chan bool)
	go func() {
		_, _ = http.Get(url)
		ch <- true
	}()
	return ch
}
