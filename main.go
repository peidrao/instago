package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/peidrao/instago/src/database"
)

func main() {
	godotenv.Load()
	dsn := "host=localhost user=" + os.Getenv("POSTGRES_USER") + " password=" + os.Getenv("POSTGRES_PASSWORD") + " dbname=" + os.Getenv("POSTGRES_DB") + " port=5432"
	database.Connect(dsn)
	database.Migrate()
}
