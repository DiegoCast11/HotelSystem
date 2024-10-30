package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql" // Driver de MySQL
	"github.com/joho/godotenv"         // Cargar variables de entorno
)

// Función para conectar a la base de datos
func ConnectDB() (*sql.DB, error) {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error al cargar el archivo .env")
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	// Verificar que la conexión sea válida
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	log.Println("Conectado a la base de datos correctamente")
	return db, nil
}
