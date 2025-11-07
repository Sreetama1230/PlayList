package handler

import (
	"Playlist/db"
	"Playlist/models"
	"net/http"

	"strconv"

	"github.com/gin-gonic/gin"
)

type songRequest struct {
	Title  string `json:"title"`
	Artist string `json:"artist"`
}
type request struct {
	Name        string        `json:"name"`
	Description string        `json:"description"`
	SongRequest []songRequest `json:"songs"`
}

func CreatePlaylist(c *gin.Context) {
	var input request
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request playload"})
		return
	}
	var songs []models.Song
	for _, s := range input.SongRequest {
		rs := models.Song{
			Title:  s.Title,
			Artist: s.Artist,
		}
		songs = append(songs, rs)
	}
	p := models.Playlist{
		Name:        input.Name,
		Description: input.Description,
		Songs:       songs,
	}

	if err := db.DB.Create(&p).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to save the playlist"})
		return
	}
	c.JSON(http.StatusCreated, p)
}

func GetPlaylist(c *gin.Context) {
	providedId := c.Param("id")
	id, err := strconv.Atoi(providedId)
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "provided id is invalid"})
		return
	}
	var playlist models.Playlist

	if err := db.DB.Model(&models.Playlist{}).Preload("Songs").Find(&playlist, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to fetch the playlist"})
		return
	}
	c.JSON(http.StatusOK, playlist)
}
