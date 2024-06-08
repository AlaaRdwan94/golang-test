package main

import (
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
)

var (
	limiterMap = make(map[string]*ratelimit.Bucket)
	mu         sync.Mutex
)

func main() {
	r := gin.Default()

	// Define rate limiting middleware
	r.Use(rateLimitMiddleware())

	// Define your POST endpoint
	r.POST("/api/rate-limit", handlePostRequest)

	r.Run(":8080")
}

func rateLimitMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()

		mu.Lock()
		bucket, ok := limiterMap[ip]
		if !ok {
			// Create a new rate limiter for this IP address
			bucket = ratelimit.NewBucketWithQuantum(
				// Specify rate limit (e.g., 100 requests per minute)
				100, 100, 1000,
			)
			limiterMap[ip] = bucket
		}
		mu.Unlock()

		if bucket.TakeAvailable(1) < 1 {
			// Rate limit exceeded, return 429 Too Many Requests
			c.AbortWithStatus(http.StatusTooManyRequests)
			return
		}

		c.Next()
	}
}

func handlePostRequest(c *gin.Context) {
	// Handle POST request logic here
	c.JSON(http.StatusOK, gin.H{
		"message": "POST request successful",
	})
}
