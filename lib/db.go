package lib

import (
	"errors"
	"log"
	"os"

	"taskflow/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// ConnectToDb connects to the postgresql database
func ConnectToDb() {
	var err error
	dbUrl := os.Getenv("DATABASE_URL") // defined in the docker-compose.yml file
	if dbUrl == "" {
		log.Fatal(errors.New("DATABASE_URL environment variable is not set"))
	}

	dsn := "host=postgres user=user password=password dbname=taskflow port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	log.Println("Connected to database")
}

// Migrate migrates the models
func Migrate() {
	DB.AutoMigrate(&models.User{}, &models.TodoItem{})
}
