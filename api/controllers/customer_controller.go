package controllers

import (
	"Hotelsystem/api/models"
	"Hotelsystem/firebase"
	"Hotelsystem/internal/database"
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// CreateCustomer maneja la creación de un cliente.
// CreateCustomer maneja la creación de un cliente.
func CreateCustomer(w http.ResponseWriter, r *http.Request) {
	var customer models.Customer
	err := json.NewDecoder(r.Body).Decode(&customer)
	if err != nil {
		http.Error(w, "Error al decodificar el cuerpo de la solicitud", http.StatusBadRequest)
		return
	}

	// Verificar si el número de teléfono ya está registrado
	var existingPhone string
	queryCheck := "SELECT phone FROM customers WHERE phone = ?"
	err = database.DB.QueryRow(queryCheck, customer.Phone).Scan(&existingPhone)
	if err != nil && err != sql.ErrNoRows {
		http.Error(w, "Error al consultar la base de datos", http.StatusInternalServerError)
		return
	}

	if existingPhone != "" {
		http.Error(w, "Error! Teléfono ya registrado", http.StatusConflict) // 409 Conflict
		return
	}

	// Insertar cliente a la base de datos
	query := "INSERT INTO customers (name, email, phone, registrydate) VALUES (?, ?, ?, ?)"
	result, err := database.DB.Exec(query, customer.Name, customer.Email, customer.Phone, customer.RegistryDate)
	if err != nil {
		http.Error(w, "Error al insertar el cliente en la base de datos", http.StatusInternalServerError)
		return
	}

	// Obtener el ID del cliente recién insertado
	customerID, err := result.LastInsertId()
	if err != nil {
		http.Error(w, "Error al obtener el ID del cliente insertado", http.StatusInternalServerError)
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
		CustomerID int    `json:"customerId"`
		Token      string `json:"token"`
	}

	var req VerifyRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Error al decodificar el cuerpo de la solicitud", http.StatusBadRequest)
		return
	}

	// verifica token y actualiza campo en bd
	if err := firebase.VerifyPhoneToken(database.DB, req.Token, req.CustomerID); err != nil {
		http.Error(w, "Error al verificar el token de teléfono", http.StatusInternalServerError)
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

	var customer models.Customer
	query := "SELECT customerid, name, email, phone, registrydate, phone_verified FROM customers WHERE phone = ?"
	err := database.DB.QueryRow(query, phone).Scan(&customer.CustomerID, &customer.Name, &customer.Email, &customer.Phone, &customer.RegistryDate, &customer.Phone_verified)
	if err != nil {
		http.Error(w, "Error al consultar la base de datos", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(customer)
}
