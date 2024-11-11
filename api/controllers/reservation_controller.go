package controllers

import (
	"Hotelsystem/api/models"
	"Hotelsystem/internal/repository"
	"Hotelsystem/pkg/contextkey"
	"encoding/json"
	"log"
	"net/http"
	"time"
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

func CreateReservation(w http.ResponseWriter, r *http.Request) {
	// Obtener userId y phone desde el contexto (middleware de autenticación)
	userID, ok := r.Context().Value(contextkey.UserIDKey).(int)

	// Verificar si los valores fueron correctamente extraídos
	if !ok {
		http.Error(w, "Token inválido", http.StatusUnauthorized)
		return
	}

	// Decodificar solicitud
	var req models.ReservationRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Error en el formato de los datos de la reserva", http.StatusBadRequest)
		return
	}

	// Validar que el teléfono esté verificado
	isVerified, err := repository.IsPhoneVerified(userID)
	if err != nil {
		http.Error(w, "Error al verificar el estado del teléfono", http.StatusInternalServerError)
		return
	}
	if !isVerified {
		http.Error(w, "Por favor verifica tu número de teléfono antes de hacer una reserva", http.StatusForbidden)
		return
	}

	// Verificar si el usuario ya tiene reservas pendientes
	pendingCount, err := repository.CountPendingReservations(userID)
	if err != nil {
		http.Error(w, "Error al verificar reservas pendientes", http.StatusInternalServerError)
		return
	}
	if pendingCount >= 2 {
		http.Error(w, "No puedes realizar más de dos reservas sin confirmar", http.StatusForbidden)
		return
	}

	// Convertir las fechas de check-in y check-out
	checkInDate, err := time.Parse("2006-01-02", req.CheckIn)
	if err != nil {
		http.Error(w, "Fecha de check-in inválida", http.StatusBadRequest)
		return
	}
	checkOutDate, err := time.Parse("2006-01-02", req.CheckOut)
	if err != nil {
		http.Error(w, "Fecha de check-out inválida", http.StatusBadRequest)
		return
	}

	// Calcular el total de noches de estancia
	nights := int(checkOutDate.Sub(checkInDate).Hours() / 24)
	if nights <= 0 {
		http.Error(w, "La fecha de check-out debe ser posterior a la de check-in", http.StatusBadRequest)
		return
	}

	// Verificar disponibilidad de habitación y obtener RoomID
	roomID, err := repository.GetAvailableRoomID(req.RoomType, req.CheckIn, req.CheckOut)
	if err != nil {
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}

	// Obtener el precio de la habitación
	roomPrice, err := repository.GetRoomPrice(req.RoomType)
	log.Print(roomPrice)
	if err != nil {
		http.Error(w, "Error al obtener el precio de la habitación", http.StatusInternalServerError)
		return
	}

	// Calcular el monto total
	totalAmount := float64(nights) * roomPrice

	// Crear la reserva
	reservation := models.Reservation{
		CustomerID:      userID,
		RoomID:          roomID,
		ReservationDate: time.Now().Format("2006-01-02"),
		CheckIn:         req.CheckIn,
		CheckOut:        req.CheckOut,
		State:           0, // Estado 0 para "pendiente"
		TotalAmount:     totalAmount,
	}

	if err := repository.CreateReservation(&reservation); err != nil {
		http.Error(w, "Error al crear la reserva", http.StatusInternalServerError)
		return
	}

	// Respuesta exitosa
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message":     "Reserva creada exitosamente",
		"reservation": reservation,
	})
}
