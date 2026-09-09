package main

import (
	"fmt"
	"sync"
	"time"
)

type bucket struct {
	tokens int
	mu     sync.RWMutex
}

// ComplexRateLimiter implements a basic token bucket algorithm.
// It supports 'Allow' being called from multiple goroutines by multiple
// clients.
type ComplexRateLimiter struct {
	mu sync.RWMutex
	// tokenBuckets keeps just a pointer to a bucket so we can work with it
	// without the need of putting it back into the map at then end of each
	// operation.
	tokenBuckets   map[string]*bucket
	maxTokens      int
	refillDuration time.Duration

	ticker    *time.Ticker
	done      chan bool
	closeOnce sync.Once
}

func NewComplexRateLimiter(maxTokens int, refillDuration time.Duration) *ComplexRateLimiter {
	srl := &ComplexRateLimiter{
		mu:             sync.RWMutex{},
		tokenBuckets:   make(map[string]*bucket),
		maxTokens:      maxTokens,
		refillDuration: refillDuration,
	}

	srl.refill()

	return srl
}

// refill spawns a goroutine that fills the rate limiter bucket
// with one token at the configured refill duration.
func (srl *ComplexRateLimiter) refill() {
	srl.ticker = time.NewTicker(srl.refillDuration)
	srl.done = make(chan bool)

	go func() {
		for {
			select {
			case <-srl.ticker.C:
				// only read lock the map
				srl.mu.RLock()

				for _, bucket := range srl.tokenBuckets {
					// lock the actual bucket we want to add more tokens
					bucket.mu.Lock()
					if bucket.tokens < srl.maxTokens {
						bucket.tokens++
					}
					bucket.mu.Unlock()
				}

				srl.mu.RUnlock()

			case <-srl.done:
				srl.ticker.Stop()
				return
			}
		}
	}()
}

// Allow returns true if the caller has enough tokens to be allowed to call
// another service.
// Each caller is identified by the key it provides and get's its own tokens.
// If Allow is called with an unkown key a new bucket will be opened, starting
// with the configured maxTokens.
func (srl *ComplexRateLimiter) Allow(key string) bool {
	// if there is no key provided simple return false
	if len(key) == 0 {
		return false
	}

	// rlock only for the check if the key already exists in the store.
	srl.mu.RLock()
	foundBucket, found := srl.tokenBuckets[key]
	srl.mu.RUnlock()

	if !found {
		// only if we couldn't find the key do a full lock of the bucket!
		srl.mu.Lock()
		// do another lookup if the bucket wasn't inserted from a concurrent
		// routine.
		// if still not, then insert it and release the lock afterwards.
		foundBucket, found = srl.tokenBuckets[key]
		if !found {
			newBucket := &bucket{tokens: srl.maxTokens, mu: sync.RWMutex{}}
			foundBucket = newBucket
			srl.tokenBuckets[key] = newBucket
		}
		srl.mu.Unlock()
	}

	foundBucket.mu.Lock()
	defer foundBucket.mu.Unlock()

	if foundBucket.tokens <= 0 {
		return false
	}

	foundBucket.tokens--

	return true
}

// Close must be called before exiting the application.
// It will gracefully shutdown running goroutines.
func (srl *ComplexRateLimiter) Close() {
	// make sure to only close this channel once because it will
	// panic on the second close!
	srl.closeOnce.Do(func() {
		close(srl.done)
	})
}

func (srl *ComplexRateLimiter) Debug() {
	srl.mu.RLock()
	defer srl.mu.RUnlock()
	for name, bucket := range srl.tokenBuckets {
		bucket.mu.RLock()
		fmt.Printf("[%s] tokens %d\n", name, bucket.tokens)
		bucket.mu.RUnlock()
	}
}
