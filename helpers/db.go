package helpers

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// DB adalah objek koneksi database
var DB *sql.DB

// InitDB untuk menginisialisasi koneksi database
func InitDB() {
	var err error
	// Menghubungkan ke database MySQL
	DB, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/belajar")
	if err != nil {
		log.Fatal(err)
	}
	if err := DB.Ping(); err != nil {
		log.Fatal(err)
	}
	log.Println("Database connection successful!")
}
