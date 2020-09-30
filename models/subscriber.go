package models

type Subscriber struct {
	ID        int `json:"id"`
	ChannelID int `json:"channel_id"`
	CreatorID int `json:"creator_id"`
}
