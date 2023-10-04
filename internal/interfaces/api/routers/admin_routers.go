package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/peidrao/instago/internal/domain/repository"
	"github.com/peidrao/instago/internal/interfaces/api/handler"
	"github.com/peidrao/instago/internal/interfaces/api/middlewares"
)

func setupAdminRoutes(admin *gin.RouterGroup, adminRepository *repository.AdminRepository, userRepository *repository.UserRepository) {
	adminHandler := handler.NewAdminHandler(adminRepository)
	admin.Use(middlewares.AuthMiddleware())
	admin.Use(middlewares.SetUserMiddleware(userRepository))
	admin.Use(middlewares.IsAdminUser())
	{
		admin.GET("/users", adminHandler.GetAllUsersHandler)
	}
}
