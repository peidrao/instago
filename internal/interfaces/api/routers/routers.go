package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/peidrao/instago/internal/domain/repository"
	"github.com/peidrao/instago/internal/interfaces/api/handler"
	"github.com/peidrao/instago/internal/interfaces/api/middlewares"
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
	api.Use(middlewares.SetUserMiddleware(userRepo))

	api.DELETE("users/", userHandler.RemoveUser)
	api.POST("follow/", userHandler.FollowUser)
	api.POST("unfollow/", userHandler.UnfollowUser)
	api.GET("followers/:username", userHandler.GetFollowers)
	api.GET("followings/:username", userHandler.GetFollowing)
	api.Use(middlewares.IsAdminUser())

	api.GET("users/:username", userHandler.GetUser)
	api.GET("users/", userHandler.GetAllUsers)

	return router
}
