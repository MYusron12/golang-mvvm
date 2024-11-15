package routes

import (
	"go-beckend/controllers"

	"github.com/gorilla/mux"
)

// SetupRoutes mendefinisikan semua routes
func SetupRoutes(router *mux.Router) {
	router.HandleFunc("/users", controllers.GetAllUsers).Methods("GET")
	router.HandleFunc("/users/{id:[0-9]+}", controllers.GetUserByID).Methods("GET")
	router.HandleFunc("/users", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/users/{id:[0-9]+}", controllers.UpdateUser).Methods("PUT")
	router.HandleFunc("/users/{id:[0-9]+}", controllers.DeleteUser).Methods("DELETE")
}
