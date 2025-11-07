package main

import (
	"Playlist/db"
	"Playlist/handler"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {

	db.InitDB()
	r := gin.Default()

	r.POST("/playlists", handler.CreatePlaylist)
	r.GET("/playlists/:id", handler.GetPlaylist)

	err := r.Run(":8080")
	if err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
