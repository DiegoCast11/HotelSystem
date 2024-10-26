package models

// Customer representa la estructura de un cliente.
type Customer struct {
	CustomerID     int    `json:"customerId"`
	Name           string `json:"name"`
	Email          string `json:"email"`
	Phone          string `json:"phone"`
	RegistryDate   string `json:"registryDate"`
	Phone_verified bool   `json:"phone_verified"`
}
