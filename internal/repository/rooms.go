package repository

import (
	"Hotelsystem/api/models"
	"Hotelsystem/internal/database"
	"errors"
)

// FetchRooms obtiene todas las habitaciones de la base de datos junto con las imágenes asociadas
func FetchRooms() ([]models.Room, error) {
	var rooms []models.Room
	// Consulta para obtener las habitaciones con su tipo correspondiente
	query := `
		SELECT r.roomId, r.roomName, rt.description AS description, r.roomTypeId, rt.roomtype AS roomtype
		FROM rooms r
		JOIN room_types rt ON r.roomTypeId = rt.roomTypeId
	`
	rows, err := database.DB.Query(query)
	if err != nil {
		return nil, errors.New("error al ejecutar la consulta para obtener habitaciones")
	}
	defer rows.Close()

	// Itera sobre los resultados y llena el slice de rooms
	for rows.Next() {
		var room models.Room
		err := rows.Scan(&room.RoomID, &room.RoomName, &room.Description, &room.RoomTypeID, &room.Type)
		if err != nil {
			return nil, errors.New("error al escanear los resultados de las habitaciones")
		}

		// Ahora obtenemos las imágenes asociadas con esta habitación
		imagesQuery := `
			SELECT imageUrl
			FROM room_images
			WHERE roomTypeId = ?
		`
		imageRows, err := database.DB.Query(imagesQuery, room.RoomTypeID)
		if err != nil {
			return nil, errors.New("error al ejecutar la consulta para obtener las imágenes")
		}
		defer imageRows.Close()

		var images []string
		for imageRows.Next() {
			var image string
			err := imageRows.Scan(&image)
			if err != nil {
				return nil, errors.New("error al escanear las imágenes")
			}
			images = append(images, image)
		}

		// Asignamos las imágenes a la habitación
		room.Images = images
		rooms = append(rooms, room)
	}

	// Verifica si hubo algún error al recorrer las filas
	if err := rows.Err(); err != nil {
		return nil, errors.New("error al recorrer los resultados de las habitaciones")
	}

	return rooms, nil
}
