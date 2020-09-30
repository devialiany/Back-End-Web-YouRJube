package models

type Membership struct {
	ID          int    `json:"id"`
	CreatorID   int    `json:"creator_id"`
	Type        int    `json:"type"`
	Date        string `json:"date"`
	ExpiredDate string `json:"expired_date"`
	Status      string `json:"status"`
}
