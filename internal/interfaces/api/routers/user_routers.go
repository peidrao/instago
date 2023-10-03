package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/peidrao/instago/internal/domain/repository"
	"github.com/peidrao/instago/internal/interfaces/api/handler"
)

func setupUserRoutes(users *gin.RouterGroup, userRepo *repository.UserRepository) {
	userHandler := handler.NewUserHandler(userRepo)

	users.PUT("", userHandler.UpdateUserHandler)
	users.PUT("picture/", userHandler.UpdatePictureUserHandler)
	users.GET(":username/", userHandler.GetUserHandler)
	users.GET("suggestions/", userHandler.GetSuggestionsForUserHandler)
	users.GET("me/", userHandler.UserMeHandler)
}
