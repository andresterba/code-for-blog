package main

import (
	"fmt"
	"sync"
	"time"
)

// SimpleRateLimiter implements a basic token bucket algorithm.
// It supports 'Allow' being called from multiple goroutines.
type SimpleRateLimiter struct {
	mu             sync.RWMutex
	tokens         int
	maxTokens      int
	refillDuration time.Duration

	ticker    *time.Ticker
	done      chan bool
	closeOnce sync.Once
}

func NewSimpleRateLimiter(maxTokens int, refillDuration time.Duration) *SimpleRateLimiter {
	srl := &SimpleRateLimiter{
		mu:             sync.RWMutex{},
		tokens:         maxTokens,
		maxTokens:      maxTokens,
		refillDuration: refillDuration,
	}

	srl.refill()

	return srl
}

// refill spawns a goroutine that fills the rate limiter bucket
// with one token at the configured refill duration.
func (srl *SimpleRateLimiter) refill() {
	srl.ticker = time.NewTicker(srl.refillDuration)
	srl.done = make(chan bool)

	go func() {
		for {
			select {
			case <-srl.ticker.C:
				srl.mu.Lock()
				if srl.tokens < srl.maxTokens {
					srl.tokens++
				}
				srl.mu.Unlock()

			case <-srl.done:
				srl.ticker.Stop()
				return
			}
		}
	}()
}

// Allow returns true if the caller has enough tokens to be allowed to call
// another service.
// In this simple implementation the key is ignored, as there is only a single
// bucket for all potential outgoing service calls.
func (srl *SimpleRateLimiter) Allow(key string) bool {
	srl.mu.Lock()
	defer srl.mu.Unlock()

	if srl.tokens <= 0 {
		return false
	}

	srl.tokens--

	return true
}

// Close must be called before exiting the application.
// It will gracefully shutdown running goroutines.
func (srl *SimpleRateLimiter) Close() {
	// make sure to only close this channel once because it will
	// panic on the second close!
	srl.closeOnce.Do(func() {
		close(srl.done)
	})
}

func (srl *SimpleRateLimiter) Debug() {
	srl.mu.RLock()
	defer srl.mu.RUnlock()
	fmt.Printf("tokens %d\n", srl.tokens)
}
