package api

import (
	_ "github.com/swaggo/swag" // This is needed for Swagger documentation
)

// @title           CostMasterHub API
// @version         1.0
// @description     API for managing products, materials, and colors in the CostMasterHub system
// @contact.name    API Support
// @contact.email   support@costmasterhub.com
// @license.name    Apache 2.0
// @license.url     http://www.apache.org/licenses/LICENSE-2.0.html
// @host           localhost:8080
// @BasePath       /
// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.
func SwaggerInfo() {
    // This function exists only for Swagger annotations
} 