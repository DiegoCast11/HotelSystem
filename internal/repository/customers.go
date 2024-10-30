package repository

import (
	"Hotelsystem/api/models"
	"Hotelsystem/internal/database"
	"database/sql"
	"errors"
)

func CreateCustomer(customer *models.Customer) (int64, error) {
	// Verificar si el número de teléfono ya está registrado
	var existingPhone string
	queryCheck := "SELECT phone FROM customers WHERE phone = ?"
	err := database.DB.QueryRow(queryCheck, customer.Phone).Scan(&existingPhone)
	if err != nil && err != sql.ErrNoRows {
		return 0, errors.New("error al verificar el número de teléfono")
	}
	if existingPhone != "" {
		return 0, errors.New("error! teléfono ya registrado")
	}

	// Insertar cliente en la base de datos
	query := "INSERT INTO customers (name, email, phone, registrydate) VALUES (?, ?, ?, ?)"
	result, err := database.DB.Exec(query, customer.Name, customer.Email, customer.Phone, customer.RegistryDate)
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
