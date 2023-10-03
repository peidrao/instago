package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/peidrao/instago/internal/domain/repository"
	"github.com/peidrao/instago/internal/interfaces/api/handler"
)

func setupAuthRoutes(auth *gin.RouterGroup, userRepo *repository.UserRepository) {
	userHandler := handler.NewUserHandler(userRepo)

	auth.POST("users/", userHandler.CreateUserHandler)
	auth.POST("login/", userHandler.LoginHandler)
	auth.POST("token_is_valid/", userHandler.TokenIsValidHandler)
}
