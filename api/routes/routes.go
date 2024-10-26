package routes

import (
	"Hotelsystem/api/controllers"

	"github.com/gorilla/mux"
)

// RegisterRoutes configura las rutas de la API.
func RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/rooms", controllers.GetRooms).Methods("GET")
	router.HandleFunc("/customers", controllers.CreateCustomer).Methods("POST")
	router.HandleFunc("/customers/verify-phone", controllers.VerifyCustomerPhone).Methods("POST")
	router.HandleFunc("/customers/{phone}", controllers.GetCustomerByPhone).Methods("GET")
}
