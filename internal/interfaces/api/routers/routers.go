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

	followRepository := repository.NewFollowRepository(db)
	followHandler := handler.NewFollowHandler(userRepo, followRepository)

	api := router.Group("/api")

	api.POST("users/", userHandler.CreateUser)

	api.POST("login/", userHandler.LoginHandler)

	api.Use(middlewares.AuthMiddleware())
	api.Use(middlewares.SetUserMiddleware(userRepo))
	api.GET("users/:username", userHandler.GetUser)
	api.GET("me/", userHandler.UserMe)

	api.POST("follow/", followHandler.FollowUser)
	api.POST("unfollow/", followHandler.UnfollowUser)
	api.GET("following/:username", followHandler.GetFollowing)
	api.GET("followers/:username", followHandler.GetFollowers)

	api.GET("following/requests/", followHandler.MeRequestFollowing)
	api.POST("following/cancel/", followHandler.CancelRequest)
	api.GET("followers/requests/", followHandler.GetRequestsFollowers)
	api.POST("followers/requests/", followHandler.AcceptRequest)

	api.Use(middlewares.IsAdminUser())

	api.GET("users/", userHandler.GetAllUsers)

	return router
}
