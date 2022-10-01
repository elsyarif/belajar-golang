package golang_database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

// GetConnection
/**
* TODO: step 2 Databse pooling
*
**/
func GetConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/golang_db?parseTime=true")

	if err != nil {
		panic(err.Error())
	}

	db.SetMaxIdleConns(10)                  // Pengaturan jumlah koneksi minimal yang dibuat
	db.SetMaxOpenConns(100)                 // Pengaturan jumlah koneksi maksimal yang dibuat
	db.SetConnMaxIdleTime(5 * time.Minute)  // Pengaturan berapa lama koneksi yang sidah tidak digunakan akan dihapus
	db.SetConnMaxLifetime(60 * time.Minute) // Pengaturan berapa lama koneksi boleh digunakan

	return db
}
