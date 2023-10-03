package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/peidrao/instago/internal/domain/repository"
	"github.com/peidrao/instago/internal/interfaces/api/handler"
)

func setupFeedRoutes(
	feed *gin.RouterGroup,
	userRepository *repository.UserRepository,
	feedRepository *repository.FeedRepository,
	postRepository *repository.PostRepository,
) {
	feedHandler := handler.NewFeedHandler(userRepository, postRepository, feedRepository)
	feed.GET("", feedHandler.FeedMeHandler)

}
