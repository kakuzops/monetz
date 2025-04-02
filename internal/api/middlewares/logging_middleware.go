package middlewares

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func LoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Registra o início da requisição
		startTime := time.Now()

		// Processa a requisição
		c.Next()

		// Registra o fim da requisição
		duration := time.Since(startTime)
		log.Printf(
			"Method: %s | URL: %s | Status: %d | Duration: %v",
			c.Request.Method,
			c.Request.URL.Path,
			c.Writer.Status(),
			duration,
		)
	}
}
