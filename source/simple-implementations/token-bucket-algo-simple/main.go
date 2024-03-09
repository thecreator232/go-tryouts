package main

import (
	"fmt"
	"math"
	"sync"
	"time"
)

type TokenBucket struct {
	tokens        int64
	maxTokens     int64
	refillRate    int64
	lastRefilTime time.Time
	mutex         sync.Mutex
}

func NewTokenBucket(rate int64, maxTokens int64) *TokenBucket {
	return &TokenBucket{
		tokens:        maxTokens,
		maxTokens:     maxTokens,
		lastRefilTime: time.Now(),
		refillRate:    rate,
	}
}

func (t *TokenBucket) refill() {
	now := time.Now()
	duration := time.Since(t.lastRefilTime)
	tokentoBeAdded := (int64(duration.Seconds()) * t.refillRate)
	t.tokens += int64(math.Min(float64(tokentoBeAdded), float64(t.maxTokens)))
	t.lastRefilTime = now
}

func (t *TokenBucket) IsRequestAllowed() bool {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	t.refill()

	if t.tokens > 0 {
		t.tokens -= 1
		return true
	}
	return false

}

func main() {
	tt := NewTokenBucket(3, 5)
	val := tt.IsRequestAllowed()
	fmt.Println("Request is allowed ? -> ", val)
	val = tt.IsRequestAllowed()
	fmt.Println("Request is allowed ? -> ", val)
	val = tt.IsRequestAllowed()
	fmt.Println("Request is allowed ? -> ", val)
	val = tt.IsRequestAllowed()
	fmt.Println("Request is allowed ? -> ", val)
	val = tt.IsRequestAllowed()
	fmt.Println("Request is allowed ? -> ", val)
	val = tt.IsRequestAllowed()
	fmt.Println("Request is allowed ? -> ", val)

	time.Sleep(1 * time.Second)

	val = tt.IsRequestAllowed()
	fmt.Println("Request is allowed ? -> ", val)
	val = tt.IsRequestAllowed()
	fmt.Println("Request is allowed ? -> ", val)
	val = tt.IsRequestAllowed()
	fmt.Println("Request is allowed ? -> ", val)
	val = tt.IsRequestAllowed()
	fmt.Println("Request is allowed ? -> ", val)
	val = tt.IsRequestAllowed()
	fmt.Println("Request is allowed ? -> ", val)
	val = tt.IsRequestAllowed()
	fmt.Println("Request is allowed ? -> ", val)
	val = tt.IsRequestAllowed()
	fmt.Println("Request is allowed ? -> ", val)
	val = tt.IsRequestAllowed()
	fmt.Println("Request is allowed ? -> ", val)
	val = tt.IsRequestAllowed()
	fmt.Println("Request is allowed ? -> ", val)

	time.Sleep(2 * time.Second)

	val = tt.IsRequestAllowed()
	fmt.Println("Request is allowed ? -> ", val)
	val = tt.IsRequestAllowed()
	fmt.Println("Request is allowed ? -> ", val)
	val = tt.IsRequestAllowed()
	fmt.Println("Request is allowed ? -> ", val)
	val = tt.IsRequestAllowed()
	fmt.Println("Request is allowed ? -> ", val)
	val = tt.IsRequestAllowed()
	fmt.Println("Request is allowed ? -> ", val)
	val = tt.IsRequestAllowed()
	fmt.Println("Request is allowed ? -> ", val)
	val = tt.IsRequestAllowed()
	fmt.Println("Request is allowed ? -> ", val)
	val = tt.IsRequestAllowed()
	fmt.Println("Request is allowed ? -> ", val)
	val = tt.IsRequestAllowed()
	fmt.Println("Request is allowed ? -> ", val)

}
