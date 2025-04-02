package handlers

import (
	"monetz/database"
	"monetz/internal/models"
	"monetz/internal/repositories"
	"monetz/internal/services"

	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateProduct(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Obtém o ID do usuário do contexto
	userID, _ := c.Get("user_id")
	product.UserID = userID.(string)

	productService := services.NewProductService(repositories.NewProductRepository(database.DB))
	if err := productService.CreateProduct(&product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product created successfully"})
}

func ListProducts(c *gin.Context) {
	// Obtém o ID do usuário do contexto
	userID, _ := c.Get("user_id")

	productService := services.NewProductService(repositories.NewProductRepository(database.DB))
	products, err := productService.ListProducts(userID.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list products"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"products": products})
}
