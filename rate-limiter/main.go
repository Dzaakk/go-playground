package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

type RateLimiter struct {
	rate       float64
	bucketSize int
	mu         sync.Mutex
	tokens     float64
	lastRefill time.Time
	clients    map[string]*ClientBucket
}

type ClientBucket struct {
	tokens     float64
	lastRefill time.Time
}

func NewRateLimiter(rate float64, bucketSize int) *RateLimiter {
	return &RateLimiter{
		rate:       rate,
		bucketSize: bucketSize,
		tokens:     float64(bucketSize),
		lastRefill: time.Now(),
		clients:    make(map[string]*ClientBucket),
	}
}

func (rl *RateLimiter) refillTokens() {
	now := time.Now()
	elapsed := now.Sub(rl.lastRefill).Seconds()
	rl.tokens = min(float64(rl.bucketSize), rl.tokens+(elapsed*rl.rate))
	rl.lastRefill = now
}

func (rl *RateLimiter) refillClientTokens(clientID string) {
	client, exists := rl.clients[clientID]
	now := time.Now()

	if !exists {
		// Create new client bucket with full tokens
		rl.clients[clientID] = &ClientBucket{
			tokens:     float64(rl.bucketSize),
			lastRefill: now,
		}
		return
	}

	elapsed := now.Sub(client.lastRefill).Seconds()
	client.tokens = min(float64(rl.bucketSize), client.tokens+(elapsed*rl.rate))
	client.lastRefill = now
}

func (rl *RateLimiter) Allow() bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	rl.refillTokens()

	if rl.tokens >= 1 {
		rl.tokens--
		return true
	}

	return false
}

func (rl *RateLimiter) AllowClient(clientID string) bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	rl.refillClientTokens(clientID)

	client := rl.clients[clientID]
	if client.tokens >= 1 {
		client.tokens--
		return true
	}

	return false
}

func RateLimitMiddleware(rl *RateLimiter, perClient bool) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			var allowed bool

			if perClient {

				clientID := r.RemoteAddr
				allowed = rl.AllowClient(clientID)
			} else {
				allowed = rl.Allow()
			}

			if !allowed {
				w.Header().Set("Retry-After", "1")
				http.Error(w, "Rate limit exceeded", http.StatusTooManyRequests)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

func main() {
	globalLimiter := NewRateLimiter(5, 10)

	clientLimiter := NewRateLimiter(2, 5)

	helloHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, World!")
	})

	statusHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "API is up and running")
	})

	http.Handle("/hello", RateLimitMiddleware(globalLimiter, false)(helloHandler))

	http.Handle("/status", RateLimitMiddleware(clientLimiter, true)(statusHandler))

	fmt.Println("Server starting on :8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Server error: %s\n", err)
	}
}

func min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}
