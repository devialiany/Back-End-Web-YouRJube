package models

type Video struct {
	ID             int    `json:"id"`
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
	Location       string `json:"location"`

	CategoryID int `json:"category_id"`
	CreatorID  int `json:"creator_id"`
}
