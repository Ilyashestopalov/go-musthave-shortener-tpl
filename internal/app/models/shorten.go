package models

//easyjson:json
type APIShortenURL struct {
	URL string `json:"url" binding:"required"`
}
