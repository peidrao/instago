package routers

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/peidrao/instago/internal/domain/repository"
	"github.com/peidrao/instago/internal/interfaces/api/middlewares"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	router := gin.Default()

	setupCors(router)

	router.Static("/static", "./static")

	userRepository := repository.NewUserRepository(db)
	followRepository := repository.NewFollowRepository(db)
	postRepository := repository.NewPostRepository(db)
	adminRepository := repository.NewAdminRepository(db)
	feedRepository := repository.NewFeedRepository(db)

	api := router.Group("/api")

	auth := api.Group("/auth")
	setupAuthRoutes(auth, userRepository)

	authenticated := api.Group("/")
	authenticated.Use(middlewares.AuthMiddleware())
	authenticated.Use(middlewares.SetUserMiddleware(userRepository))
	{

		users := authenticated.Group("users/")
		setupUserRoutes(users, userRepository)

		follows := authenticated.Group("follow/")
		setupFollowRoutes(follows, userRepository, followRepository)

		following := authenticated.Group("following/")
		setupFollowingRoutes(following, userRepository, followRepository)

		posts := authenticated.Group("posts/")
		setupPostRoutes(posts, userRepository, followRepository, postRepository)

		feed := authenticated.Group("feed/")
		setupFeedRoutes(feed, userRepository, feedRepository, postRepository)
	}

	admin := api.Group("/admin")
	setupAdminRoutes(admin, adminRepository, userRepository)

	return router
}

func setupCors(router *gin.Engine) {
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowCredentials = true
	config.AddAllowHeaders("authorization")
	router.Use(cors.New(config))
}
