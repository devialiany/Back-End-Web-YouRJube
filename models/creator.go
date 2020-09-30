package models

type Creator struct {
	ID               int         `json:"id"`
	Name             string      `json:"name"`
	Email            string      `json:"email"`
	JoinDate         string      `json:"join_date"`
	Description      string      `json:"description"`
	PhotoProfile     string      `json:"photo_profile"`
	PhotoBackground  string      `json:"photo_background"`
	MembershipStatus int         `json:"membership_status"`
	Subscriber       int         `json:"subscriber"`
	Restricted       int         `json:"restricted"`
	Link             string      `json:"link"`
	Videos           []*Video    `json:"videos"`
	Playlists        []*Playlist `json:"playlists`
}
