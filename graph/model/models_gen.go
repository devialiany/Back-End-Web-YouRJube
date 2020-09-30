// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Userplaylist struct {
	CreatorID  int `json:"creator_id"`
	PlaylistID int `json:"playlist_id"`
}

type NewComment struct {
	Content    string `json:"content"`
	DateUpload string `json:"date_upload"`
	VideoID    int    `json:"video_id"`
	CreatorID  int    `json:"creator_id"`
}

type NewCommunity struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	PostImage   string `json:"post_image"`
	Date        string `json:"date"`
	CreatorID   int    `json:"creator_id"`
}

type NewCreator struct {
	Name             string `json:"name"`
	Email            string `json:"email"`
	JoinDate         string `json:"join_date"`
	Description      string `json:"description"`
	PhotoProfile     string `json:"photo_profile"`
	PhotoBackground  string `json:"photo_background"`
	MembershipStatus int    `json:"membership_status"`
	Subscriber       int    `json:"subscriber"`
	Restricted       int    `json:"restricted"`
}

type NewDetailPlaylist struct {
	PlaylistID int `json:"playlist_id"`
	VideoID    int `json:"video_id"`
}

type NewLikeComment struct {
	CommentID int  `json:"comment_id"`
	LikerID   int  `json:"liker_id"`
	Status    bool `json:"status"`
}

type NewLikeCommunity struct {
	CommunityID int  `json:"community_id"`
	LikerID     int  `json:"liker_id"`
	Status      bool `json:"status"`
}

type NewLikeVideo struct {
	VideoID int  `json:"video_id"`
	LikerID int  `json:"liker_id"`
	Status  bool `json:"status"`
}

type NewMembership struct {
	CreatorID   int    `json:"creator_id"`
	Type        int    `json:"type"`
	Date        string `json:"date"`
	ExpiredDate string `json:"expired_date"`
	Status      string `json:"status"`
}

type NewPlaylist struct {
	Title       string `json:"title"`
	Date        string `json:"date"`
	Description string `json:"description"`
	Viewer      int    `json:"viewer"`
	CreatorID   int    `json:"creator_id"`
	Type        int    `json:"type"`
}

type NewReply struct {
	Content    string `json:"content"`
	DateUpload string `json:"date_upload"`
	CommentID  int    `json:"comment_id"`
	CreatorID  int    `json:"creator_id"`
}

type NewSubscriber struct {
	ChannelID int `json:"channel_id"`
	CreatorID int `json:"creator_id"`
}

type NewUserPlaylist struct {
	CreatorID  int `json:"creator_id"`
	PlaylistID int `json:"playlist_id"`
}

type NewVideo struct {
	Title          string `json:"title"`
	Description    string `json:"description"`
	Duration       int    `json:"duration"`
	DateUpload     string `json:"date_upload"`
	DatePublish    string `json:"date_publish"`
	PhotoThumbnail string `json:"photo_thumbnail"`
	VideoFile      string `json:"video_file"`
	Restricted     int    `json:"restricted"`
	Privacy        int    `json:"privacy"`
	Premium        int    `json:"premium"`
	Status         int    `json:"status"`
	Viewer         int    `json:"viewer"`
	CreatorID      int    `json:"creator_id"`
	CategoryID     int    `json:"category_id"`
	Location       string `json:"location"`
}

type UpdateBanner struct {
	PhotoBackground string `json:"photo_background"`
}

type UpdateDate struct {
	Date string `json:"date"`
}

type UpdateDescriptionPlaylist struct {
	Description string `json:"description"`
}

type UpdateIcon struct {
	PhotoProfile string `json:"photo_profile"`
}

type UpdateLikeComment struct {
	Status bool `json:"status"`
}

type UpdateLikeCommunity struct {
	Status bool `json:"status"`
}

type UpdateLikeVideo struct {
	Status bool `json:"status"`
}

type UpdateMembershipStatusCreator struct {
	MembershipStatus int `json:"membership_status"`
}

type UpdateTitlePlaylist struct {
	Title string `json:"title"`
}

type UpdateTypePlaylist struct {
	Type int `json:"type"`
}

type UpdateVideo struct {
	Title          string `json:"title"`
	Description    string `json:"description"`
	PhotoThumbnail string `json:"photo_thumbnail"`
	Privacy        int    `json:"privacy"`
}