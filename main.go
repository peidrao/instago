package main

import (
	"github.com/joho/godotenv"
	"github.com/peidrao/instago/src/database"
	"github.com/peidrao/instago/src/routers"
)

func main() {
	godotenv.Load()

	db := database.Init()

	router := routers.SetupRouter(db)
	router.Run(":8080")
}
