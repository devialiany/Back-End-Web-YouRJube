package models

type Feedbackcommunity struct {
	ID          string `json:"id"`
	CommunityID int    `json:"community_id"`
	LikerID     int    `json:"liker_id"`
	Status      bool   `json:"status"`
}
