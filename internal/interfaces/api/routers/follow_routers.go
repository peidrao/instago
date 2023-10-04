package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/peidrao/instago/internal/domain/repository"
	"github.com/peidrao/instago/internal/interfaces/api/handler"
)

func setupFollowRoutes(
	follows *gin.RouterGroup,
	userRepo *repository.UserRepository,
	followRepository *repository.FollowRepository,
) {
	followHandler := handler.NewFollowHandler(userRepo, followRepository)

	follows.POST("", followHandler.FollowUserHandler)
	follows.POST("delete/", followHandler.UnfollowUserHandler)
	follows.GET(":username/", followHandler.GetFollowersHandler)
	follows.GET("requests/", followHandler.GetFollowersRequestHandler)
	follows.POST("requests/", followHandler.AcceptRequestHandler)
}
