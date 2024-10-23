package models

type Room struct {
	RoomID      string `json:"roomId"`
	RoomName    string `json:"roomName"`
	Type        string `json:"type"`
	Description string `json:"description"`
	Capacity    int    `json:"capacity"`
	Dimensions  int    `json:"dimensions"`
}
