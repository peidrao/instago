package main

import (
	"github.com/joho/godotenv"
	"github.com/peidrao/instago/internal/infrastructure"
	"github.com/peidrao/instago/internal/interfaces/api/routers"
)

func main() {
	godotenv.Load()

	db := infrastructure.Init()

	router := routers.SetupRouter(db)
	router.Run(":8080")
}
