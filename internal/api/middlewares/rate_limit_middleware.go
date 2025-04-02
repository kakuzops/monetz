package middlewares

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

type rateLimiter struct {
	ips map[string]int
	mu  sync.Mutex
}

var (
	limiter = rateLimiter{ips: make(map[string]int)}
	limit   = 100             // Número máximo de requisições
	window  = 1 * time.Minute // Janela de tempo para o rate limiting
)

func RateLimitingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()

		limiter.mu.Lock()
		defer limiter.mu.Unlock()

		// Verifica se o IP excedeu o limite
		if limiter.ips[ip] >= limit {
			c.JSON(http.StatusTooManyRequests, gin.H{"error": "Too many requests"})
			c.Abort()
			return
		}

		// Incrementa o contador de requisições para o IP
		limiter.ips[ip]++

		// Reseta o contador após a janela de tempo
		go func() {
			time.Sleep(window)
			limiter.mu.Lock()
			limiter.ips[ip]--
			limiter.mu.Unlock()
		}()

		c.Next()
	}
}
