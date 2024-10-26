package database

import (
	"database/sql"
)

var DB *sql.DB

// SetDB establece la conexión de la base de datos a nivel global
func SetDB(database *sql.DB) {
	DB = database
}
