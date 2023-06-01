package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/peidrao/instago/src/controllers"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	api := router.Group("/api")

	api.POST("user/register", controllers.RegisterUser)
	return router
}
