package models

type Community struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	PostImage   string `json:"post_image"`
	Date        string `json:"date"`
	CreatorID   int    `json:"creator_id"`
}
