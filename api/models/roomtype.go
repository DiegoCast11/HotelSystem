package models

type RoomType struct {
	RoomTypeId  int    `json:"roomTypeId"`
	Description string `json:"description"`
	Capacity    int    `json:"capacity"`
	Dimensions  int    `json:"dimensions"`
}
