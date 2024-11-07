package server

import (
	"Hotelsystem/api/routes"
	"database/sql"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
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

	// Configurar CORS
	corsHandler := configureCORS()

	// Aplica el middleware de CORS al enrutador
	s.router.Use(corsHandler.Handler)
	return s
}

// Configuración de CORS
func configureCORS() *cors.Cors {
	return cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000", "https://sf-project-one.vercel.app"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})
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
