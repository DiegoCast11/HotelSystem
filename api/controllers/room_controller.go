package controllers

import (
	"Hotelsystem/internal/repository"
	"encoding/json"
	"net/http"
)

func GetRooms(w http.ResponseWriter, r *http.Request) {
	// Llamada a la función FetchRooms para obtener las habitaciones
	rooms, err := repository.FetchRooms()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Devuelve la respuesta como JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(rooms)
}
