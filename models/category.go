package models


type Category struct {
	ID       	int          	`json:"id"`
	Category 	string          `json:"category"`
	Videos   	[]*Video 		`json:"videos"`
}
