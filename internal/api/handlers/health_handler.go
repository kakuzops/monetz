package handlers

import (
	"github.com/gin-gonic/gin"
	"time"
)

// @Summary     Health check endpoint
// @Description Get service health status
// @Tags        health
// @Produce     json
// @Success     200 {object} map[string]interface{}
// @Router      / [get]
func Health(c *gin.Context) {
	healthStatus := gin.H{
		"status":    "healthy",
		"timestamp": time.Now().UTC(),
		"service":   "Vasco da gama e nada mais",
	}
	
	c.JSON(200, healthStatus)
}

