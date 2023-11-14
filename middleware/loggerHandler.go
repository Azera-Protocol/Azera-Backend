package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer
		start := time.Now()
		path := c.Request.URL.Path

		// Process request
		c.Next()

		// Calculate latency
		latency := time.Since(start)

		// Get status code
		status := c.Writer.Status()

		// Log format
		log.Printf("%s %s - %d %v", c.Request.Method, path, status, latency)

		// Check if there were any errors, and log them
		if len(c.Errors) > 0 {
			log.Printf("Request errors: %v", c.Errors)
		}
	}
}
