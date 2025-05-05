package handlers

import (
	"monetz/database"
	"monetz/internal/models"
	"monetz/internal/repositories"
	"monetz/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateSeller(c *gin.Context) {
	var seller models.Seller
	if err := c.ShouldBindJSON(&seller); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, _ := c.Get("user_id")
	seller.UserID = userID.(string)

	sellerService := services.NewSellerService(repositories.NewSellerRepository(database.DB))
	if err := sellerService.CreateSeller(&seller); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create seller"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Seller created successfully"})
}

func ListSellers(c *gin.Context) {
	userID, _ := c.Get("user_id")

	sellerService := services.NewSellerService(repositories.NewSellerRepository(database.DB))
	sellers, err := sellerService.ListSellers(userID.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list sellers"})
		return
	}
	c.JSON(http.StatusOK, sellers)
}

func DeleteSeller(c *gin.Context) {
	userID, _ := c.Get("user_id")
	sellerID := c.Param("id")

	sellerService := services.NewSellerService(repositories.NewSellerRepository(database.DB))
	if err := sellerService.DeleteSeller(sellerID, userID.(string)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete seller"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Seller deleted successfully"})
}

func UpdateSeller(c *gin.Context) {
	userID, _ := c.Get("user_id")
	sellerID := c.Param("id")

	var seller models.Seller
	if err := c.ShouldBindJSON(&seller); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sellerService := services.NewSellerService(repositories.NewSellerRepository(database.DB))
	if err := sellerService.UpdateSeller(sellerID, userID.(string), &seller); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update seller"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Seller updated successfully"})
}
