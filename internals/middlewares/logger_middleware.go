package middlewares

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

// Types and DI

type ILoggerMiddleware interface {
	Initialize() gin.HandlerFunc
}

type LoggerMiddleware struct {
}

func NewLoggerMiddleware() ILoggerMiddleware {
	return &LoggerMiddleware{}
}

// Functions

func (m *LoggerMiddleware) Initialize() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		// before request
		c.Next()
		// after request

		latency := time.Since(t)
		log.Printf("latency: %s ms", latency)
	}
}
