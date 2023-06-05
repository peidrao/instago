package database

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/peidrao/instago/src/domain/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	godotenv.Load()

	dsn := "host=localhost user=" + os.Getenv("POSTGRES_USER") + " password=" + os.Getenv("POSTGRES_PASSWORD") + " dbname=" + os.Getenv("POSTGRES_DB") + " port=5432"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	log.Println("Connected to database")

	db.AutoMigrate(&models.User{})

	return db
}
