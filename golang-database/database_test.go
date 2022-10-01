package golang_database

import (
	"database/sql"                     // ^1
	_ "github.com/go-sql-driver/mysql" // ^2
	"testing"
)

//  TODO: step 1, import database driver *^
func TestOpenConnection(t *testing.T) {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/golang_db")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
}
