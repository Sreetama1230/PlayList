package models

type Playlist struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Songs       []Song `json:"songs" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type Song struct {
	ID         uint   `json:"id" gorm:"primaryKey"`
	Title      string `json:"title"`
	Artist     string `json:"artist"`
	PlaylistID uint   `json:"playlistId"`
}
