package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/peidrao/instago/internal/domain/repository"
	"github.com/peidrao/instago/internal/interfaces/api/handler"
)

func setupPostRoutes(
	posts *gin.RouterGroup,
	userRepository *repository.UserRepository,
	followRepository *repository.FollowRepository,
	postRepository *repository.PostRepository,
) {
	postHandler := handler.NewPostHandler(userRepository, postRepository, followRepository)

	posts.POST("", postHandler.CreatePostHandler)
	posts.GET("me/", postHandler.GetMePostsHandler)
	posts.GET(":id/", postHandler.GetPostHandler)
	posts.DELETE(":id/", postHandler.DeletePostHandler)
}
