package handlers

import (
	"log"
	"monetz/database"
	"monetz/internal/models"
	"monetz/internal/repositories"
	"monetz/internal/services"
	"monetz/internal/utils"
	"strconv"

	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// @Summary     Register new user
// @Description Register a new user in the system
// @Tags        auth
// @Accept      json
// @Produce     json
// @Param       user body models.User true "User registration details"
// @Success     200 {object} map[string]interface{}
// @Failure     400 {object} map[string]interface{}
// @Router      /api/register [post]
func Register(c *gin.Context) {
	var userInput models.User
	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validação adicional da senha
	if len(userInput.PasswordHash) < 6 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Password must be at least 6 characters long"})
		return
	}

	// Gera o hash da senha com custo específico
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userInput.PasswordHash), 12)
	if err != nil {
		log.Printf("Error hashing password: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to process password"})
		return
	}

	// Verifica se o hash foi gerado corretamente
	if len(hashedPassword) < 60 {
		log.Printf("Generated hash is too short: %d bytes", len(hashedPassword))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate secure password"})
		return
	}

	user := &models.User{
		Name:         userInput.Name,
		Email:        userInput.Email,
		PasswordHash: string(hashedPassword),
		Role:         userInput.Role,
	}

	userRepo := repositories.NewUserRepository(database.DB)
	if err := userRepo.Create(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

// @Summary     Login user
// @Description Authenticate user and get JWT token
// @Tags        auth
// @Accept      json
// @Produce     json
// @Param       credentials body object true "User credentials"
// @Success     200 {object} map[string]interface{}
// @Failure     401 {object} map[string]interface{}
// @Router      /api/login [post]
func Login(c *gin.Context) {
	var input struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		log.Printf("Bind error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userService := services.NewUserService(repositories.NewUserRepository(database.DB))
	user, err := userService.FindUserByEmail(input.Email)
	if err != nil {
		log.Printf("User not found: %s", input.Email)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	if len(user.PasswordHash) < 60 {
		log.Printf("Stored hash is invalid for user %s", input.Email)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid password hash stored"})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(input.Password))
	if err != nil {
		log.Printf("Password mismatch for user %s: %v", input.Email, err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	userID := strconv.FormatUint(uint64(user.ID), 10)
	token, err := utils.GenerateToken(userID, user.Role)
	if err != nil {
		log.Printf("Token generation error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"name":    user.Name,
		"email":   user.Email,
		"role":    user.Role,
		"token":   token,
	})
}
