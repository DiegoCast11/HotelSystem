package routes

import (
	"Hotelsystem/api/controllers"
	"Hotelsystem/api/middleware"

	"github.com/gorilla/mux"
)

// RegisterRoutes configura las rutas de la API.
func RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/rooms", controllers.GetRooms).Methods("GET")
	router.HandleFunc("/customers", controllers.CreateCustomer).Methods("POST")
	router.HandleFunc("/customers/verify-phone", controllers.VerifyCustomerPhone).Methods("POST")
	router.HandleFunc("/customers/{phone}", controllers.GetCustomerByPhone).Methods("GET")
	router.HandleFunc("/login", controllers.Login).Methods("POST")
	router.HandleFunc("/availability", controllers.CheckAvailability).Methods("GET")

	// Crear subenrutador para rutas privadas
	privateRouter := router.PathPrefix("/").Subrouter()
	privateRouter.Use(middleware.AuthMiddleware)

	// Rutas privadas
	privateRouter.HandleFunc("/reservations", controllers.CreateReservation).Methods("POST")
}
