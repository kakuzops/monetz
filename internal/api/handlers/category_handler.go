package handlers

import (
	"monetz/database"
	"monetz/internal/models"
	"monetz/internal/repositories"
	"monetz/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateCategory godoc
// @Summary Create a new category
// @Description Create a new category for the authenticated user
// @Tags categories
// @Accept json
// @Produce json
// @Security Bearer
// @Param category body models.Category true "Category object"
// @Success 201 {object} models.Category
// @Failure 400 {object} map[string]string
// @Router /categories [post]
func CreateCategory(c *gin.Context) {
	categoryService := services.NewCategoryService(repositories.NewCategoryRepository(database.DB))
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	var category models.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	category.UserID = userID.(string)

	if err := categoryService.CreateCategory(&category); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create category"})
		return
	}

	c.JSON(http.StatusCreated, category)
}

// GetCategories godoc
// @Summary List all categories
// @Description Get all categories for the authenticated user
// @Tags categories
// @Produce json
// @Security Bearer
// @Success 200 {array} models.Category
// @Router /categories [get]
func GetCategories(c *gin.Context) {
	categoryService := services.NewCategoryService(repositories.NewCategoryRepository(database.DB))
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	categories, err := categoryService.ListCategories(userID.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch categories"})
		return
	}

	c.JSON(http.StatusOK, categories)
}

// GetCategory godoc
// @Summary Get a category by ID
// @Description Get a specific category by ID for the authenticated user
// @Tags categories
// @Produce json
// @Security Bearer
// @Param id path int true "Category ID"
// @Success 200 {object} models.Category
// @Failure 404 {object} map[string]string
// @Router /categories/{id} [get]
func GetCategory(c *gin.Context) {
	categoryService := services.NewCategoryService(repositories.NewCategoryRepository(database.DB))
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category ID"})
		return
	}

	category, err := categoryService.GetCategory(uint(id), userID.(string))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	}

	c.JSON(http.StatusOK, category)
}

// UpdateCategory godoc
// @Summary Update a category
// @Description Update a specific category for the authenticated user
// @Tags categories
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path int true "Category ID"
// @Param category body models.Category true "Category object"
// @Success 200 {object} models.Category
// @Failure 404 {object} map[string]string
// @Router /categories/{id} [put]
func UpdateCategory(c *gin.Context) {
	categoryService := services.NewCategoryService(repositories.NewCategoryRepository(database.DB))
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category ID"})
		return
	}

	var category models.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	category.ID = uint(id)
	category.UserID = userID.(string)

	if err := categoryService.UpdateCategory(&category); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	}

	c.JSON(http.StatusOK, category)
}

// DeleteCategory godoc
// @Summary Delete a category
// @Description Delete a specific category for the authenticated user
// @Tags categories
// @Security Bearer
// @Param id path int true "Category ID"
// @Success 204 "No Content"
// @Failure 404 {object} map[string]string
// @Router /categories/{id} [delete]
func DeleteCategory(c *gin.Context) {
	categoryService := services.NewCategoryService(repositories.NewCategoryRepository(database.DB))
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category ID"})
		return
	}

	if err := categoryService.DeleteCategory(uint(id), userID.(string)); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	}

	c.Status(http.StatusNoContent)
}
