package repository

import (
	"Hotelsystem/api/models"
	"Hotelsystem/internal/database"
	"errors"
)

// CheckAvailability devuelve las reservas de habitaciones

func CheckAvailability() ([]models.ReservationResponse, error) {
	var reservations []models.ReservationResponse

	query := `
		SELECT r.reservationId, r.roomid, r.checkin, r.checkout, rt.roomtype
		FROM reservations r
		JOIN rooms ro ON r.roomid = ro.roomId
    JOIN room_types rt ON ro.roomTypeId = rt.roomTypeId
		WHERE r.checkin >= CURDATE();

	`
	row, err := database.DB.Query(query)

	if err != nil {
		return nil, errors.New("error al obtener las reservas")
	}
	defer row.Close()

	for row.Next() {
		var reservation models.ReservationResponse
		err := row.Scan(&reservation.ReservationID, &reservation.RoomID, &reservation.CheckIn, &reservation.CheckOut, &reservation.RoomType)
		if err != nil {
			return nil, errors.New("error al escanear las reservas")
		}
		reservations = append(reservations, reservation)
	}

	if err := row.Err(); err != nil {
		return nil, errors.New("error al obtener las reservas")
	}

	return reservations, nil

}
