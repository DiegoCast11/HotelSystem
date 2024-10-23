package server

import (
	"Hotelsystem/api/controllers"
	"database/sql"

	"github.com/gorilla/mux"
)

// Server representa la estructura del servidor.
type Server struct {
	router *mux.Router
	db     *sql.DB
}

// NewServer crea una nueva instancia del servidor.
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
	// Aquí registrarías las rutas
	controllers.SetDB(s.db)
	s.router.HandleFunc("/rooms", controllers.GetRooms).Methods("GET")
}

// Obtener el enrutador
func (s *Server) Router() *mux.Router {
	return s.router
}
