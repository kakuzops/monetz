package api

import (
	_ "monetz/docs" // swagger docs
	"monetz/internal/api/handlers"
	"monetz/internal/api/middlewares"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// More flexible CORS configuration for development
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:*", "http://127.0.0.1:*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization", "X-Requested-With"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowWildcard:    true,
		MaxAge:           12 * time.Hour,
	}))

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
		authGroup.DELETE("/products/:id", handlers.DeleteProduct)
		authGroup.PUT("/products/:id", handlers.UpdateProduct)
		authGroup.GET("/products/count", handlers.GetCountProduct)

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

		// Seller routes
		authGroup.POST("/sellers", handlers.CreateSeller)
		authGroup.GET("/sellers", handlers.ListSellers)
		authGroup.DELETE("/sellers/:id", handlers.DeleteSeller)
		authGroup.PUT("/sellers/:id", handlers.UpdateSeller)

		dashboardGroup := authGroup.Group("/admin")
		dashboardGroup.Use(middlewares.AdminMiddleware())
		{
			dashboardGroup.GET("/users", handlers.ListUsers)
			dashboardGroup.GET("/users/:id", handlers.GetUserDetails)
		}
	}

	return router
}
