package config

import (
	"go-crud/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectionDataBase() {
	dsn := "host=localhost user=postgres password=admin dbname=postgres port=5432 sslmode=disable"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	DB = database
}

func InitDatabase() {
	err := DB.AutoMigrate(&models.Item{})
	if err != nil {
		log.Fatalf("Error migration database %v", err)
	}
	log.Println("Database Succesfully migrated")
}
