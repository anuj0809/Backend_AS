package database

import (
	"os"

	"github.com/anuj0809/Backend_AS/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"fmt"
)

// setup the database connection
var DB *gorm.DB

func ConnectToDB() {
	var err error

	// err = godotenv.Load()
	// if err != nil {
	// 	panic("Failed to load .env file")
	// }

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))
	// os.Getenv("DB")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = DB.AutoMigrate(&models.Players{})
	if err != nil {
		panic(err)
	}

}
