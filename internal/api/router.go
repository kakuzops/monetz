package api

import (
	_ "monetz/docs" // swagger docs
	"monetz/internal/api/handlers"
	"monetz/internal/api/middlewares"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Swagger documentation endpoint
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Middlewares globais
	router.Use(middlewares.LoggingMiddleware())
	router.Use(middlewares.RateLimitingMiddleware())

	// Rotas públicas
	router.POST("/api/register", handlers.Register)
	router.POST("/api/login", handlers.Login)
	router.GET("/", handlers.Health)

	// Rotas protegidas (requerem autenticação)
	authGroup := router.Group("/api")
	authGroup.Use(middlewares.JWTMiddleware())
	{
		// Product routes
		authGroup.POST("/products", handlers.CreateProduct)
		authGroup.GET("/products", handlers.ListProducts)

		// Material routes
		authGroup.POST("/materials", handlers.CreateMaterial)
		authGroup.GET("/materials", handlers.ListMaterials)

		// Color routes
		authGroup.POST("/colors", handlers.CreateColor)
		authGroup.GET("/colors", handlers.ListColors)

		// Category routes
		authGroup.POST("/categories", handlers.CreateCategory)
		authGroup.GET("/categories", handlers.GetCategories)
		authGroup.GET("/categories/:id", handlers.GetCategory)
		authGroup.PUT("/categories/:id", handlers.UpdateCategory)
		authGroup.DELETE("/categories/:id", handlers.DeleteCategory)

		dashboardGroup := authGroup.Group("/admin")
		dashboardGroup.Use(middlewares.AdminMiddleware())
		{
			dashboardGroup.GET("/users", handlers.ListUsers)
			dashboardGroup.GET("/users/:id", handlers.GetUserDetails)
		}
	}

	return router
}
