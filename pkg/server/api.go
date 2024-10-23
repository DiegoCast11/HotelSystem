package server

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"Hotelsystem/pkg/models"

	"github.com/gorilla/mux"
)

type Server struct {
	router *mux.Router
	db     *sql.DB
}

// Nueva instancia del servidor
func NewServer(db *sql.DB) *Server {
	s := &Server{
		router: mux.NewRouter(),
		db:     db,
	}
	s.routes() // Definir rutas
	return s
}

// Manejador de rutas
func (s *Server) routes() {
	s.router.HandleFunc("/rooms", s.getRooms).Methods("GET")
}

// Manejador para obtener habitaciones
func (s *Server) getRooms(w http.ResponseWriter, r *http.Request) {
	rows, err := s.db.Query("SELECT roomId, roomName, type, description, capacity, dimensions FROM rooms")
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

// Obtener el enrutador
func (s *Server) Router() http.Handler {
	return s.router
}
