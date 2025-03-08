package database

import (
	"fmt"
	"log"
	"monetz/src/config/models"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

const defaultPort = 5432

func loadEnv() error {
	envFiles := []string{".env.local", ".env", ".env.development"}

	for _, file := range envFiles {
		if err := godotenv.Load(file); err == nil {
			fmt.Printf("Using environment file: %s\n", file)
			return nil
		}
	}

	return fmt.Errorf("no environment file found. Tried: %v", envFiles)
}

func InitDB() {
	if err := loadEnv(); err != nil {
		log.Fatal(err)
	}

	portStr := os.Getenv("DB_PORT")
	port := defaultPort
	if portStr != "" {
		var err error
		port, err = strconv.Atoi(portStr)
		if err != nil {
			log.Fatalf("Erro ao converter porta do banco de dados: %v", err)
		}
	}

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"), port, os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}

	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatalf("Erro ao obter conexão SQL: %v", err)
	}

	err = sqlDB.Ping()
	if err != nil {
		log.Fatalf("Erro ao pingar o banco de dados: %v", err)
	}

	// AutoMigrate will create tables based on your models
	err = AutoMigrate()
	if err != nil {
		log.Fatalf("Erro ao realizar migração: %v", err)
	}

	fmt.Println("Conexão com o banco de dados estabelecida com sucesso!")
}

// AutoMigrate performs database migrations for all models
func AutoMigrate() error {
	// Add your models here to auto-migrate
	// Example: err := DB.AutoMigrate(&User{}, &Product{}, &Order{})
	err := DB.AutoMigrate(
		&models.User{},
	// Add your models here
	// &models.User{},
	// &models.Product{},
	)

	if err != nil {
		return fmt.Errorf("erro ao realizar migração: %v", err)
	}

	return nil
}
