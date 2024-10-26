package server

import (
	"Hotelsystem/api/routes"
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
	// Conectar a la base de datos
	s := &Server{
		router: mux.NewRouter(),
		db:     db,
	}
	s.routes() // Definir rutas
	return s
}

// Manejador de rutas
func (s *Server) routes() {
	// Registrar rutas del controlador
	routes.RegisterRoutes(s.router)
}

// Obtener el enrutador
func (s *Server) Router() *mux.Router {
	return s.router
}
