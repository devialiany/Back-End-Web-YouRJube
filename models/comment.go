package models

type Comment struct {
	ID         int    `json:"id"`
	Content    string `json:"content"`
	DateUpload string `json:"date_upload"`
	VideoID    int    `json:"video_id"`
	CreatorID  int    `json:"creator_id"`

	Replies []*Reply `json:"replies"`
}
