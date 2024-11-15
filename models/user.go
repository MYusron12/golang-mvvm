package models

import (
	"database/sql"

	"go-beckend/helpers"

	_ "github.com/go-sql-driver/mysql" // Driver MySQL
)

// User adalah representasi dari tabel user di database
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// CreateUser menambahkan user baru ke database
func CreateUser(user *User) error {
	db := helpers.GetDB()
	_, err := db.Exec("INSERT INTO users (username, email, password) VALUES (?, ?, ?)", user.Username, user.Email, user.Password)
	return err
}

// GetAllUsers mengambil semua user dari database
func GetAllUsers() ([]User, error) {
	var users []User
	db := helpers.GetDB()

	rows, err := db.Query("SELECT id, username, email, password FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.Password); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

// GetUserByID mengambil user berdasarkan ID
func GetUserByID(id int) (*User, error) {
	var user User
	db := helpers.GetDB()

	err := db.QueryRow("SELECT id, username, email, password FROM users WHERE id = ?", id).Scan(&user.ID, &user.Username, &user.Email, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

// UpdateUser memperbarui data user di database
func UpdateUser(user *User) error {
	db := helpers.GetDB()

	_, err := db.Exec("UPDATE users SET username = ?, email = ?, password = ? WHERE id = ?", user.Username, user.Email, user.Password, user.ID)
	return err
}

// DeleteUser menghapus user berdasarkan ID
func DeleteUser(id int) error {
	db := helpers.GetDB()

	_, err := db.Exec("DELETE FROM users WHERE id = ?", id)
	return err
}
