package models

type Feedbackcomment struct {
	ID        string `json:"id"`
	CommentID int    `json:"comment_id"`
	LikerID   int    `json:"liker_id"`
	Status    bool   `json:"status"`
}
