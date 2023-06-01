package main

import (
	"github.com/joho/godotenv"
	"github.com/peidrao/instago/src/database"
	"github.com/peidrao/instago/src/routers"
)

func main() {
	godotenv.Load()
	database.ConnectDB()
	router := routers.SetupRouter()
	router.Run(":8080")
}
