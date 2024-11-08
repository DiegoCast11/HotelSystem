package controllers

import (
	"Hotelsystem/api/models"
	"Hotelsystem/internal/repository"
	"Hotelsystem/services"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// CreateCustomer maneja la creación de un cliente.
func CreateCustomer(w http.ResponseWriter, r *http.Request) {
	var customer models.Customer
	err := json.NewDecoder(r.Body).Decode(&customer)
	if err != nil {
		http.Error(w, "Error al decodificar el cuerpo de la solicitud", http.StatusBadRequest)
		return
	}

	hash, err := services.HashPassword(customer.Password)
	if err != nil {
		http.Error(w, "error al encriptar la contraseña", http.StatusInternalServerError)
		return
	}
	customer.Password = hash

	// Llamar al repositorio para crear el cliente
	customerID, err := repository.CreateCustomer(&customer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Respuesta exitosa con el código 201 y el ID del cliente creado
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message":    "Customer created successfully",
		"customerId": customerID,
	})
}

// VerifyCustomerPhone verifica el token del teléfono y actualiza la base de datos.
func VerifyCustomerPhone(w http.ResponseWriter, r *http.Request) {
	type VerifyRequest struct {
		Phone string `json:"Phone"`
	}

	var req VerifyRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Error al decodificar el cuerpo de la solicitud", http.StatusBadRequest)
		return
	}

	err := repository.UpdatePhoneVerification(req.Phone)

	if err != nil {
		http.Error(w, "Error al verificar el teléfono", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Phone verified successfully",
	})
}

// GetCustomerByPhone obtiene un cliente por su número de teléfono.
func GetCustomerByPhone(w http.ResponseWriter, r *http.Request) {
	phone := mux.Vars(r)["phone"]

	customer, err := repository.GetCustomerByPhone(phone)
	if err != nil {
		http.Error(w, "Error al obtener el cliente", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(customer)
}
