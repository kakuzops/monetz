package main

import (
	"log"
	"monetz/database"
	_ "monetz/docs"
	"monetz/internal/api"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

func init() {
	// Try to load .env file from different possible locations
	envFiles := []string{".env", ".env.dev", "../.env", "../../.env"}

	for _, file := range envFiles {
		if _, err := os.Stat(file); err == nil {
			if err := godotenv.Load(file); err != nil {
				log.Printf("Warning: Error loading %s file: %v", file, err)
			} else {
				log.Printf("Loaded environment from %s", file)
				break
			}
		}
	}
}

func main() {
	// Ensure we're in the right working directory
	if err := os.Chdir(filepath.Dir(os.Args[0])); err != nil {
		log.Printf("Warning: Failed to change working directory: %v", err)
	}

	database.InitDB()
	router := api.SetupRouter()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
