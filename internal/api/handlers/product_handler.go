package handlers

import (
	"monetz/database"
	"monetz/internal/models"
	"monetz/internal/repositories"
	"monetz/internal/services"

	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateProduct(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

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
	userID, _ := c.Get("user_id")

	productService := services.NewProductService(repositories.NewProductRepository(database.DB))
	products, err := productService.ListProducts(userID.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list products"})
		return
	}

	c.JSON(http.StatusOK, products)
}

func DeleteProduct(c *gin.Context) {
	productID := c.Param("id")

	productService := services.NewProductService(repositories.NewProductRepository(database.DB))
	if err := productService.DeleteProduct(productID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}

func UpdateProduct(c *gin.Context) {
	productID := c.Param("id")
	productService := services.NewProductService(repositories.NewProductRepository(database.DB))

	var input struct {
		Name  string  `json:"name"`
		Price float64 `json:"price"`
		Stock int     `json:"stock"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := productService.UpdateProduct(productID, input.Name, fmt.Sprintf("%.2f", input.Price), input.Stock); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product updated successfully"})
}

func GetCountProduct(c *gin.Context) {
	userID, _ := c.Get("user_id")

	productService := services.NewProductService(repositories.NewProductRepository(database.DB))
	count, err := productService.GetCountProduct(userID.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get count product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"count": count})
}
