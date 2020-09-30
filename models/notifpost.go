package models

type Notifpost struct {
	ID          int `json:"id"`
	CommunityID int `json:"community_id"`
	CreatorID   int `json:"creator_id"`
}
