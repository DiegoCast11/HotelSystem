package repository

import (
	"Hotelsystem/api/models"
	"Hotelsystem/internal/database"
	"database/sql"
	"errors"
	"fmt"
)

func CreateCustomer(customer *models.Customer) (int64, error) {

	var existingPhone string
	queryCheck := "SELECT phone FROM customers WHERE phone = ?"
	err := database.DB.QueryRow(queryCheck, customer.Phone).Scan(&existingPhone)
	if err != nil && err != sql.ErrNoRows {
		return 0, errors.New("error al verificar el número de teléfono")
	}
	if existingPhone != "" {
		return 0, errors.New("error! teléfono ya registrado")
	}

	// Check if the email is already registered
	var existingEmail string
	queryCheckEmail := "SELECT email FROM customers WHERE email = ?"
	err = database.DB.QueryRow(queryCheckEmail, customer.Email).Scan(&existingEmail)
	if err != nil && err != sql.ErrNoRows {
		return 0, errors.New("error al verificar el correo electrónico")
	}
	if existingEmail != "" {
		return 0, errors.New("error! correo electrónico ya registrado")
	}

	// Insertar cliente en la base de datos
	query := "INSERT INTO customers (name, email, phone, password) VALUES (?, ?, ?, ?)"
	result, err := database.DB.Exec(query, customer.Name, customer.Email, customer.Phone, customer.Password)
	if err != nil {
		return 0, errors.New("error al insertar el cliente en la base de datos")
	}

	// Obtener el ID del cliente recién insertado
	customerID, err := result.LastInsertId()
	if err != nil {
		return 0, errors.New("error al obtener el ID del cliente insertado")
	}

	return customerID, nil
}

func UpdatePhoneVerification(phone string) error {
	query := "UPDATE customers SET phone_verified = 1 WHERE phone = ?"
	result, err := database.DB.Exec(query, phone)
	if err != nil {
		return errors.New("error al actualizar la verificación del teléfono")
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return errors.New("error al obtener el número de filas afectadas")
	}
	if rowsAffected == 0 {
		return errors.New("error! teléfono no encontrado")
	}

	return nil
}

// GetCustomerByPhone obtiene un cliente por su número de teléfono
func GetCustomerByPhone(phone string) (*models.Customer, error) {
	var customer models.Customer
	query := "SELECT customerid, name, email, phone, registrydate, phone_verified FROM customers WHERE phone = ?"
	err := database.DB.QueryRow(query, phone).Scan(
		&customer.CustomerID,
		&customer.Name,
		&customer.Email,
		&customer.Phone,
		&customer.RegistryDate,
		&customer.Phone_verified,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("usuario no encontrado")
		}
		return nil, errors.New("error al consultar la base de datos")
	}

	return &customer, nil
}

func GetUserByEmail(email string) (models.User, error) {
	var user models.User
	query := `SELECT customerId, email, password FROM customers WHERE email = ?`
	err := database.DB.QueryRow(query, email).Scan(&user.UserID, &user.Email, &user.Password)
	if err != nil {
		return models.User{}, errors.New("usuario no encontrado")
	}
	return user, nil
}

func IsPhoneVerified(userID int) (bool, error) {
	var phoneVerified bool
	query := "SELECT phone_verified FROM customers WHERE customerId = ?"
	err := database.DB.QueryRow(query, userID).Scan(&phoneVerified)
	if err == sql.ErrNoRows {
		return false, fmt.Errorf("usuario no encontrado")
	} else if err != nil {
		return false, err
	}
	return phoneVerified, nil
}
