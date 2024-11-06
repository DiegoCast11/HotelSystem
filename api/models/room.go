package models

type Room struct {
	RoomID      string   `json:"roomId"`
	RoomName    string   `json:"roomName"`
	Type        string   `json:"type"`
	Description string   `json:"description"`
	RoomTypeID  int      `json:"roomTypeId"`
	Images      []string `json:"images"` // Lista de URLs de las imágenes asociadas
}
