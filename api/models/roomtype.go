package models

type RoomType struct {
	RoomTypeId  int      `json:"roomTypeId"`
	RoomType    string   `json:"roomtype"`
	Description string   `json:"description"`
	Capacity    int      `json:"capacity"`
	Dimensions  int      `json:"dimensions"`
	Price       int      `json:"price"`
	Images      []string `json:"images"` // Lista de URLs de las imágenes asociadas
}
