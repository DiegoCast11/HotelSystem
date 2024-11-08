package controllers

import (
	"Hotelsystem/internal/repository"
	"encoding/json"
	"net/http"
)

func CheckAvailability(w http.ResponseWriter, r *http.Request) {
	reservations, err := repository.CheckAvailability()
	if err != nil {
		http.Error(w, "error al obtener las reservaciones", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(reservations)
}
