package main

import (
	"go-beckend/helpers"
	"go-beckend/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Inisialisasi database
	helpers.InitDB()

	// Membuat router baru
	router := mux.NewRouter()

	// Setup routes
	routes.SetupRoutes(router)

	// Jalankan server
	log.Fatal(http.ListenAndServe(":8080", router))
}
