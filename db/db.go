package db

import (
	"Playlist/models"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	// dsn := "sreetama:password@tcp(127.0.0.1:3306)/playlist?charset=utf8mb4&parseTime=True&loc=Local"
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", dbUser, dbPass, dbHost, dbName)

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to open the db connection: %v", err)
	}

	if err := DB.AutoMigrate(&models.Playlist{}, &models.Song{}); err != nil {
		log.Fatalf("migrate error :%v", err)
	}
}
