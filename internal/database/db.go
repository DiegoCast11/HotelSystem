package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql" // Driver de MySQL
)

// Función para conectar a la base de datos
func ConnectDB() (*sql.DB, error) {
	// Obtén la URL de conexión de JawsDB desde el entorno
	dsn := os.Getenv("JAWSDB_URL")
	if dsn == "" {
		log.Fatal("La variable de entorno JAWSDB_URL no está configurada")
	}
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
