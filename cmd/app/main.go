package main

import (
	"log"
	"net/http"
	"os"

	"Hotelsystem/api/routes"        // Configura las rutas
	"Hotelsystem/internal/database" // Conectar a la BD
	"Hotelsystem/pkg/server"        // Inicializa el servidor
)

func main() {
	// Conectar a la base de datos
	db, err := database.ConnectDB()
	if err != nil {
		log.Fatalf("Error al conectar a la base de datos: %v", err)
	}
	defer db.Close()

	database.SetDB(db)

	// Inicializar servidor
	s := server.NewServer(db)

	// Configurar rutas
	routes.RegisterRoutes(s.Router())

	// Obtener el puerto del entorno
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Iniciar servidor HTTP
	log.Printf("Servidor corriendo en :%s", port)
	log.Fatal(http.ListenAndServe(":"+port, s.Router()))
}
