package main

import (
	"log"
	"net/http"

	"Hotelsystem/pkg/server" // Inicializa el servidor

	"Hotelsystem/internal/database" // Conectar a la BD
)

func main() {
	// Conectar a la base de datos
	db, err := database.ConnectDB()
	if err != nil {
		log.Fatalf("Error al conectar a la base de datos: %v", err)
	}
	defer db.Close()

	// Inicializar servidor
	s := server.NewServer(db)

	// Iniciar servidor HTTP
	log.Println("Servidor corriendo en :8080")
	log.Fatal(http.ListenAndServe(":8080", s.Router()))
}
