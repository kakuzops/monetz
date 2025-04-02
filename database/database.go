package database

import (
	"monetz/internal/config"
	"monetz/internal/models"

	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error
	dsn := config.GetDBConnectionString()
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to database: %v", err))
	}

	// Migrar modelos
	err = DB.AutoMigrate(&models.User{}, &models.Product{}, &models.Material{}, &models.Color{})
	if err != nil {
		panic(fmt.Sprintf("Failed to migrate database: %v", err))
	}

	fmt.Println("Successfully connected to the database!")
}
