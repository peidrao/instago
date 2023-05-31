package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/peidrao/instago/src/controllers"
	"github.com/peidrao/instago/src/database"
)

func connectDB() {
	dsn := "host=localhost user=" + os.Getenv("POSTGRES_USER") + " password=" + os.Getenv("POSTGRES_PASSWORD") + " dbname=" + os.Getenv("POSTGRES_DB") + " port=5432"
	database.Connect(dsn)
	database.Migrate()

}

func main() {
	godotenv.Load()
	connectDB()
	router := initRouter()
	router.Run(":8080")
}

func initRouter() *gin.Engine {
	router := gin.Default()
	api := router.Group("/api")

	api.POST("user/register", controllers.RegisterUser)
	return router
}
