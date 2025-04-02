package handlers

import (
	"monetz/database"
	"monetz/internal/models"
	"monetz/internal/repositories"
	"monetz/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

var validate = validator.New()

func CreateColor(c *gin.Context) {
	var color models.Color
	if err := c.ShouldBindJSON(&color); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := validate.Struct(color); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	colorService := services.NewColorService(repositories.NewColorRepository(database.DB))
	if err := colorService.CreateColor(&color); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create color"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Color created successfully"})
}

func ListColors(c *gin.Context) {
	colorService := services.NewColorService(repositories.NewColorRepository(database.DB))
	colors, err := colorService.ListColors()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list colors"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"colors": colors})
}
