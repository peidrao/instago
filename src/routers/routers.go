package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/peidrao/instago/src/handler"
	"github.com/peidrao/instago/src/internal/delivery/middlewares"
	"github.com/peidrao/instago/src/repository"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	router := gin.Default()

	userRepo := repository.NewUserRepository(db)
	userHandler := handler.NewUserHandler(userRepo)

	api := router.Group("/api")

	api.POST("users/", userHandler.RegisterUser)
	api.POST("login/", userHandler.LoginHandler)

	api.Use(middlewares.AuthMiddleware())

	api.GET("users/", userHandler.GetAllUsers)
	api.GET("users/:id", userHandler.GetUser)
	api.DELETE("users/:id", userHandler.RemoveUser)

	return router
}
