package models

// Request de reserva
type ReservationRequest struct {
	RoomType int    `json:"roomType"`
	CheckIn  string `json:"checkin"`
	CheckOut string `json:"checkout"`
}
