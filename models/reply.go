package models

type Reply struct {
	ID         int    `json:"id"`
	Content    string `json:"content"`
	DateUpload string `json:"date_upload"`
	CommentID  int    `json:"comment_id"`
	CreatorID  int    `json:"creator_id"`
}
