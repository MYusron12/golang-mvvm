package controllers

import (
	"encoding/json"
	"go-beckend/models"
	"go-beckend/views"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// CreateUser menangani request untuk membuat user baru
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		views.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if err := models.CreateUser(&user); err != nil {
		views.RespondWithError(w, http.StatusInternalServerError, "Could not create user")
		return
	}

	views.RespondWithJSON(w, http.StatusCreated, user)
}

// GetAllUsers menangani request untuk mendapatkan semua user
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := models.GetAllUsers()
	if err != nil {
		views.RespondWithError(w, http.StatusInternalServerError, "Could not fetch users")
		return
	}

	views.RespondWithJSON(w, http.StatusOK, users)
}

// GetUserByID menangani request untuk mendapatkan user berdasarkan ID
func GetUserByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	user, err := models.GetUserByID(id)
	if err != nil {
		views.RespondWithError(w, http.StatusInternalServerError, "Could not fetch user")
		return
	}

	if user == nil {
		views.RespondWithError(w, http.StatusNotFound, "User not found")
		return
	}

	views.RespondWithJSON(w, http.StatusOK, user)
}

// UpdateUser menangani request untuk memperbarui user
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	var user models.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		views.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	user.ID = id

	if err := models.UpdateUser(&user); err != nil {
		views.RespondWithError(w, http.StatusInternalServerError, "Could not update user")
		return
	}

	views.RespondWithJSON(w, http.StatusOK, user)
}

// DeleteUser menangani request untuk menghapus user
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	if err := models.DeleteUser(id); err != nil {
		views.RespondWithError(w, http.StatusInternalServerError, "Could not delete user")
		return
	}

	views.RespondWithJSON(w, http.StatusNoContent, nil)
}
