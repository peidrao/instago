package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/peidrao/instago/internal/infrastructure"
	"github.com/peidrao/instago/internal/interfaces/api/routers"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db := infrastructure.Init()

	router := routers.SetupRouter(db)
	err = router.Run(":8080")

	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
