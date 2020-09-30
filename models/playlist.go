package models

type Playlist struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Date        string `json:"date"`
	Description string `json:"description"`
	Viewer      int    `json:"viewer"`
	CreatorID   int    `json:"creator_id"`
	Type        int    `json:"type"`
}
