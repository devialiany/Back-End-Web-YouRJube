package models

type Feedbackvideo struct {
	ID      string `json:"id"`
	VideoID int    `json:"video_id"`
	LikerID int    `json:"liker_id"`
	Status  bool   `json:"status"`
}
