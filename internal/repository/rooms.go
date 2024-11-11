package repository

import (
	"Hotelsystem/api/models"
	"Hotelsystem/internal/database"
	"database/sql"
	"errors"
	"fmt"
)

// FetchRooms obtiene todas las habitaciones de la base de datos junto con las imágenes asociadas
func FetchRooms() ([]models.RoomType, error) {
	var rooms []models.RoomType
	// Consulta para obtener las habitaciones con su tipo correspondiente
	query := `	
		SELECT
			r.roomTypeId, r.roomtype, r.description, r.capacity, r.dimensions, c.price
		FROM room_types r
		JOIN costs c ON r.roomTypeId = c.roomTypeId
		WHERE
			c.startDate <= CURDATE() AND
			(c.lastDate >= CURDATE() OR c.lastDate IS NULL)
		GROUP BY r.roomTypeId;
	`
	rows, err := database.DB.Query(query)
	if err != nil {
		return nil, errors.New("error al ejecutar la consulta para obtener habitaciones")
	}
	defer rows.Close()

	// Itera sobre los resultados y llena el slice de rooms
	for rows.Next() {
		var room models.RoomType
		err := rows.Scan(&room.RoomTypeId, &room.RoomType, &room.Description, &room.Capacity, &room.Dimensions, &room.Price)
		if err != nil {
			return nil, errors.New("error al escanear los resultados de las habitaciones")
		}

		// Ahora obtenemos las imágenes asociadas con esta habitación
		imagesQuery := `
			SELECT imageUrl
			FROM room_images
			WHERE roomTypeId = ?
		`
		imageRows, err := database.DB.Query(imagesQuery, room.RoomTypeId)
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

func GetAvailableRoomID(roomType int, checkIn string, checkOut string) (string, error) {
	var roomID string
	query := `
		SELECT r.roomId
		FROM rooms r
		LEFT JOIN reservations b ON r.roomId = b.roomid
		AND b.state IN (0, 1)
		AND (
			(b.checkin <= ? AND b.checkout >= ?) OR  -- La reserva existente se superpone en el checkin
			(b.checkin <= ? AND b.checkout >= ?) OR  -- La reserva existente se superpone en el checkout
			(b.checkin >= ? AND b.checkout <= ?)     -- La reserva existente está dentro del rango deseado
		)
		WHERE r.roomTypeId = ?
		AND b.roomid IS NULL -- Asegurarse de que no haya reservas en las fechas solicitadas
		LIMIT 1
	`

	err := database.DB.QueryRow(query, checkIn, checkIn, checkOut, checkOut, checkIn, checkOut, roomType).Scan(&roomID)
	if err == sql.ErrNoRows {
		return "", fmt.Errorf("no hay habitaciones disponibles para el tipo de habitación %d en las fechas especificadas", roomType)
	} else if err != nil {
		return "", err
	}

	return roomID, nil
}

// Obtiene el precio de la habitación basada en el tipo
func GetRoomPrice(roomType int) (float64, error) {
	var price float64
	query := `SELECT price
			  FROM costs c
			  JOIN room_types r ON r.roomTypeId = c.roomTypeId
			  WHERE
				c.roomTypeId = ? AND
				startDate <= CURDATE() AND
				(lastDate >= CURDATE() OR lastDate IS NULL)
			GROUP BY r.roomTypeId;`

	err := database.DB.QueryRow(query, roomType).Scan(&price)
	if err == sql.ErrNoRows {
		return 0, fmt.Errorf("no se encontró precio para el tipo de habitación especificado")
	} else if err != nil {
		return 0, err
	}

	return price, nil
}
