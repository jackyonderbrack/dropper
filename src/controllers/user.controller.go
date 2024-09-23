package controllers

import (
	"backend/src/services"
	"encoding/json"
	"net/http"

	"gorm.io/gorm"
)

func CheckIfAnyUserExist(db *gorm.DB) http.HandlerFunc {
	userService := services.NewuserService(db)

	return func(w http.ResponseWriter, r *http.Request) {
		isAnyUser, err := userService.IsAnyUser()
		if err != nil {
			http.Error(w, "Błąd serwera", http.StatusInternalServerError)
			return
		}

		response := map[string]bool{"isAnyUser": isAnyUser}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}