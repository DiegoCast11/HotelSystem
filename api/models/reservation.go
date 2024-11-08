package models

type ReservationResponse struct {
	ReservationID int    `json:"reservationId"`
	RoomID        string `json:"roomid"`
	CheckIn       string `json:"checkin"`
	CheckOut      string `json:"checkout"`
	RoomType      string `json:"roomtype"`
}
