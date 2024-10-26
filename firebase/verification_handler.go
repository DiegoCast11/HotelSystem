package firebase

import (
	"context"
	"database/sql"
	"log"
)

// VerifyPhoneToken verifica el token de firebase y actualiza campo en bd.
func VerifyPhoneToken(db *sql.DB, token string, customerID int) error {
	app, err := FirebaseApp()
	if err != nil {
		return err
	}

	client, err := app.Auth(context.Background())

	//verificar token
	if _, err := client.VerifyIDToken(context.Background(), token); err != nil {
		log.Printf("Error al verificar el token de teléfono: %v\n", err)
		return err
	}

	// token verificado
	if err := updatePhoneVerifiedStatus(db, customerID); err != nil {
		return err
	}

	log.Printf("Usuario con ID %v ha verificado su número de teléfono\n", customerID)
	return nil
}

func updatePhoneVerifiedStatus(db *sql.DB, customerID int) error {
	query := "UPDATE customers SET phone_verified = 1 WHERE customerId = ?"
	_, err := db.Exec(query, customerID)
	if err != nil {
		log.Printf("Error al actualizar el estado de verificación del teléfono: %v\n", err)
		return err
	}
	return nil
}
