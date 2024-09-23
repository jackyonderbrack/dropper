package routes

import (
	"backend/src/controllers"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func RegisterUserRoutes(r *mux.Router, db *gorm.DB) {
	r.HandleFunc("/api/isanyuser", controllers.CheckIfAnyUserExist(db)).Methods(http.MethodGet)
}