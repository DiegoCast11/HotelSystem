package models

// Customer representa la estructura de un cliente.
type User struct {
	UserID   int    `json:"customerId"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
