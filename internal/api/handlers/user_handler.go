package handlers

import (
	"monetz/database"
	"monetz/internal/repositories"
	"monetz/internal/services"

	"github.com/gin-gonic/gin"

	"net/http"
	"strconv"
)

// @Summary     List Users
// @Description Get service List Users
// @Tags        ListUsers
// @Produce     json
// @Success     200 {object} map[string]interface{}
// @Router      / [get]
func ListUsers(c *gin.Context) {
	userService := services.NewUserService(repositories.NewUserRepository(database.DB))
	users, err := userService.ListUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list users"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"users": users})
}

func GetUserDetails(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	userService := services.NewUserService(repositories.NewUserRepository(database.DB))
	user, err := userService.GetUserByID(uint(userID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}
