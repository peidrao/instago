package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/peidrao/instago/internal/domain/repository"
	"github.com/peidrao/instago/internal/interfaces/api/handler"
)

func setupFollowingRoutes(
	following *gin.RouterGroup,
	userRepo *repository.UserRepository,
	followRepository *repository.FollowRepository,
) {
	followHandler := handler.NewFollowHandler(userRepo, followRepository)

	following.GET(":username/", followHandler.GetFollowingHandler)
	following.GET("requests/", followHandler.GetFollowingRequestHandler)
	following.POST("delete/", followHandler.CancelRequestHandler)
}
