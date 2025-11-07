package db

import (
	"Playlist/models"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := "sreetama:password@tcp(127.0.0.1:3306)/playlist?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to open the db connection: %v", err)
	}

	if err := DB.AutoMigrate(&models.Playlist{}, &models.Song{}); err != nil {
		log.Fatalf("migrate error :%v", err)
	}
}
