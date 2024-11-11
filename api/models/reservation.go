package models

type ReservationResponse struct {
	ReservationID int    `json:"reservationId"`
	RoomID        string `json:"roomid"`
	CheckIn       string `json:"checkin"`
	CheckOut      string `json:"checkout"`
	RoomType      string `json:"roomtype"`
}

type Reservation struct {
	ReservationID   int     `json:"reservationId"`
	CustomerID      int     `json:"customerid"`
	RoomID          string  `json:"roomid"`
	ReservationDate string  `json:"reservationdate"`
	CheckIn         string  `json:"checkin"`
	CheckOut        string  `json:"checkout"`
	State           int     `json:"state"`
	TotalAmount     float64 `json:"totalamount"`
}
