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
	router.Static("/static", "./static")

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
		auth.POST("users/", userHandler.CreateUserHandler)
		auth.POST("login/", userHandler.LoginHandler)
		auth.POST("token_is_valid/", userHandler.TokenIsValidHandler)
	}

	authenticated := api.Group("/")
	authenticated.Use(middlewares.AuthMiddleware())
	authenticated.Use(middlewares.SetUserMiddleware(userRepo))
	{

		users := authenticated.Group("users/")
		{
			users.PUT("", userHandler.UpdateUserHandler)
			users.PUT("picture/", userHandler.UpdatePictureUserHandler)
			users.GET(":username/", userHandler.GetUserHandler)
			users.GET("suggestions/", userHandler.GetSuggestionsForUserHandler)

			users.GET("me/", userHandler.UserMeHandler)
		}

		follows := authenticated.Group("follow/")
		{
			follows.POST("", followHandler.FollowUserHandler)
			follows.POST("delete/", followHandler.UnfollowUserHandler)
			follows.GET(":username/", followHandler.GetFollowersHandler)
			follows.GET("requests/", followHandler.GetFollowersRequestHandler)
			follows.POST("requests/", followHandler.AcceptRequestHandler)
		}

		following := authenticated.Group("following/")
		{
			following.GET(":username/", followHandler.GetFollowingHandler)
			following.GET("requests/", followHandler.GetFollowingRequestHandler)
			following.POST("delete/", followHandler.CancelRequestHandler)
		}

		posts := authenticated.Group("posts/")
		{
			posts.POST("", postHandler.CreatePostHandler)
			posts.GET("me/", postHandler.GetMePostsHandler)
			posts.GET(":id/", postHandler.GetPostHandler)
			posts.DELETE(":id/", postHandler.DeletePostHandler)
		}

		feed := authenticated.Group("feed/")
		{
			feed.GET("", feedHandler.FeedMeHandler)
		}
	}

	admin := api.Group("/admin")
	admin.Use(middlewares.AuthMiddleware())
	admin.Use(middlewares.SetUserMiddleware(userRepo))
	admin.Use(middlewares.IsAdminUser())
	{
		admin.GET("/users", userHandler.GetAllUsersHandler)

	}

	return router
}
