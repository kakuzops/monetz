package main

import (
	"log"
	"net/http"

	"monetz/src/app/user/handler"
	"monetz/src/config/database"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func healthCheck(c *gin.Context) {
	sqlDB, err := database.DB.DB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Erro ao obter conexão SQL",
		})
		return
	}

	err = sqlDB.Ping()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Erro ao conectar ao banco de dados",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"message": "Everything is fine! o....o",
	})
}

func main() {
	database.InitDB()
	sqlDB, err := database.DB.DB()
	if err != nil {
		log.Fatalf("Erro ao obter conexão SQL: %v", err)
	}
	defer sqlDB.Close()

	r := gin.Default()

	r.GET("/health", healthCheck)

	user := r.Group("/user")
	{
		user.POST("/", handler.CreateUserHandler)
		user.GET("/", handler.GetUserHandler)
		user.GET("/:id", handler.GetUserHandler)
	}

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}
