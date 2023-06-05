package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/peidrao/instago/src/handler"
	"github.com/peidrao/instago/src/repository"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	router := gin.Default()

	userRepo := repository.NewUserRepository(db)
	userHandler := handler.NewUserHandler(userRepo)

	api := router.Group("/api")

	api.POST("user/", userHandler.RegisterUser)
	api.GET("user/:id", userHandler.GetUser)

	return router
}
