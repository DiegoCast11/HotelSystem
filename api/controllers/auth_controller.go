package controllers

import (
	"Hotelsystem/internal/repository"
	"Hotelsystem/services"
	"encoding/json"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var credentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		http.Error(w, "error al decodificar credenciales", http.StatusBadRequest)
		return
	}

	user, err := repository.GetUserByEmail(credentials.Email)

	if err != nil {
		http.Error(w, "usuario no encontrado", http.StatusUnauthorized)
		return
	}

	// Verificar la contraseña ingresada con la almacenada
	if !services.CheckPassword(credentials.Password, user.Password) {
		http.Error(w, "Usuario o contraseña incorrectos", http.StatusUnauthorized)
		return
	}

	token, err := services.GenerateToken(user.UserID, user.Email)
	if err != nil {
		http.Error(w, "error al generar token", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"token": token})
}
