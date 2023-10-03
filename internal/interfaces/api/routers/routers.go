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
	followRepository := repository.NewFollowRepository(db)
	postRepository := repository.NewPostRepository(db)

	feedRepository := repository.NewFeedRepository(db)

	adminHandler := handler.NewAdminHandler(userRepo)

	api := router.Group("/api")

	auth := api.Group("/auth")
	setupAuthRoutes(auth, userRepo)

	authenticated := api.Group("/")
	authenticated.Use(middlewares.AuthMiddleware())
	authenticated.Use(middlewares.SetUserMiddleware(userRepo))
	{

		users := authenticated.Group("users/")
		setupUserRoutes(users, userRepo)

		follows := authenticated.Group("follow/")
		setupFollowRoutes(follows, userRepo, followRepository)

		following := authenticated.Group("following/")
		setupFollowingRoutes(following, userRepo, followRepository)

		posts := authenticated.Group("posts/")
		setupPostRoutes(posts, userRepo, followRepository, postRepository)

		feed := authenticated.Group("feed/")
		setupFeedRoutes(feed, userRepo, feedRepository, postRepository)
	}

	admin := api.Group("/admin")
	admin.Use(middlewares.AuthMiddleware())
	admin.Use(middlewares.SetUserMiddleware(userRepo))
	admin.Use(middlewares.IsAdminUser())
	{
		admin.GET("/users", adminHandler.GetAllUsersHandler)

	}

	return router
}
