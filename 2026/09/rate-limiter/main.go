package main

import (
	"fmt"
	"sync"
	"time"
)

type RateLimiter interface {
	Allow(key string) bool
}

type DebugLimiter interface {
	RateLimiter
	Debug()
}

func main() {
	fmt.Println("Hello World")

	// srl := NewSimpleRateLimiter(10, 2*time.Second)
	srl := NewComplexRateLimiter(10, 2*time.Second)

	// simpleDemo(srl)
	concurrentDemo(srl)

	srl.Close()
}

func simpleDemo(rl DebugLimiter) {
	for range 10 {
		rl.Allow("test")

	}

	for range 10 {
		rl.Debug()
		if rl.Allow("test") {
			fmt.Println("allowed to request")
		} else {
			fmt.Println("rate limited")
		}

		time.Sleep(time.Second)
	}
}

func concurrentDemo(rl RateLimiter) {
	var wg sync.WaitGroup

	callers := []string{"a", "b", "c"}
	// must be more than maxTokens to get some actually limited
	const requestsPerCaller = 15

	for _, caller := range callers {
		for i := range requestsPerCaller {
			wg.Add(1)
			go func(key string, n int) {
				defer wg.Done()

				if rl.Allow(key) {
					fmt.Printf("[%s] request %d: allowed to request\n", key, n)
				} else {
					fmt.Printf("[%s] request %d: rate limited\n", key, n)
				}
			}(caller, i)
		}
	}

	wg.Wait()
}
