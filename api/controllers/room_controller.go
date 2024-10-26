package controllers

import (
	"Hotelsystem/api/models"
	"Hotelsystem/internal/database"
	"encoding/json"
	"net/http"
)

// GetRooms obtiene la lista de habitaciones y la devuelve como JSON.
func GetRooms(w http.ResponseWriter, r *http.Request) {
	rows, err := database.DB.Query("SELECT roomId, roomName, type, description, capacity, dimensions FROM rooms")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var rooms []models.Room // Usa el struct Room del paquete models
	for rows.Next() {
		var room models.Room
		if err := rows.Scan(&room.RoomID, &room.RoomName, &room.Type, &room.Description, &room.Capacity, &room.Dimensions); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		rooms = append(rooms, room) // Agrega la habitación al slice
	}

	// Devuelve la respuesta como JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(rooms)
}
