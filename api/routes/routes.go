package routes

import (
	"Hotelsystem/api/controllers"

	"github.com/gorilla/mux"
)

// RegisterRoutes configura las rutas de la API.
func RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/rooms", controllers.GetRooms).Methods("GET")
}
