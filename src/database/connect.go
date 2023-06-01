package database

import "os"

func ConnectDB() {
	dsn := "host=localhost user=" + os.Getenv("POSTGRES_USER") + " password=" + os.Getenv("POSTGRES_PASSWORD") + " dbname=" + os.Getenv("POSTGRES_DB") + " port=5432"
	Connect(dsn)
	Migrate()
}
