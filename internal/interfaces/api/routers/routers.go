package routers

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/peidrao/instago/internal/domain/repository"
	"github.com/peidrao/instago/internal/interfaces/api/handler"
	"github.com/peidrao/instago/internal/interfaces/api/middlewares"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowCredentials = true
	config.AddAllowHeaders("authorization")
	router.Use(cors.New(config))

	userRepo := repository.NewUserRepository(db)
	userHandler := handler.NewUserHandler(userRepo)

	followRepository := repository.NewFollowRepository(db)
	followHandler := handler.NewFollowHandler(userRepo, followRepository)

	postRepository := repository.NewPostRepository(db)
	postHandler := handler.NewPostHandler(userRepo, postRepository, followRepository)

	feedRepository := repository.NewFeedRepository(db)
	feedHandler := handler.NewFeedHandler(userRepo, postRepository, feedRepository)

	api := router.Group("/api")

	auth := api.Group("/auth")
	{
		auth.POST("/users", userHandler.CreateUser)
		auth.POST("/login", userHandler.LoginHandler)
	}

	authenticated := api.Group("/")
	authenticated.Use(middlewares.AuthMiddleware())
	authenticated.Use(middlewares.SetUserMiddleware(userRepo))
	{

		users := authenticated.Group("users/")
		{
			users.PUT("", userHandler.UpdateUser)
			users.PUT("picture/", userHandler.UpdatePictureUser)
			users.GET(":username/", userHandler.GetUser)
			users.GET("me/", userHandler.UserMe)
		}

		follows := authenticated.Group("follow/")
		{
			follows.POST("", followHandler.FollowUser)
			follows.POST("delete/", followHandler.UnfollowUser)
			follows.GET(":username/", followHandler.GetFollowers)
			follows.GET("requests/", followHandler.GetFollowersRequest)
			follows.POST("requests/", followHandler.AcceptRequest)
		}

		following := authenticated.Group("following/")
		{
			following.GET(":username/", followHandler.GetFollowing)
			following.GET("requests/", followHandler.GetFollowingRequest)
			following.POST("delete/", followHandler.CancelRequest)
		}

		posts := authenticated.Group("posts/")
		{
			posts.POST("", postHandler.CreatePost)
			posts.GET("me/", postHandler.GetMePosts)
			posts.GET(":id/", postHandler.GetPost)
			posts.DELETE(":id/", postHandler.DeletePost)
		}

		feed := authenticated.Group("feed/")
		{
			feed.GET("", feedHandler.FeedMe)
		}
	}

	admin := api.Group("/admin")
	admin.Use(middlewares.IsAdminUser())
	{
		admin.GET("/users", userHandler.GetAllUsers)

	}

	return router
}
