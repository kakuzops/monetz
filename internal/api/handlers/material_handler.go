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

var checkValidate = validator.New()

func CreateMaterial(c *gin.Context) {
	var material models.Material
	if err := c.ShouldBindJSON(&material); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := checkValidate.Struct(material); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	materialService := services.NewMaterialService(repositories.NewMaterialRepository(database.DB))
	if err := materialService.CreateMaterial(&material); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create material"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Material created successfully"})
}

func ListMaterials(c *gin.Context) {
	materialService := services.NewMaterialService(repositories.NewMaterialRepository(database.DB))
	materials, err := materialService.ListMaterials()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list materials"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"materials": materials})
}
