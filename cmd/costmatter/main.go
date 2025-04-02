package main

import (
	"log"
	"monetz/database"
	_ "monetz/docs" // Import the docs
	"monetz/internal/api"

	"github.com/joho/godotenv"
	_ "github.com/swaggo/swag" // Required for swagger
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Warning: .env file not found: %v", err)
	}
	database.InitDB()
	router := api.SetupRouter()
	router.Run(":8080")
}
