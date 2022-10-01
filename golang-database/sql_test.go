package golang_database

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"
)

func TestExecSql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	query := "INSERT into customer(id, name) value ('al', 'Al')"
	_, err := db.ExecContext(ctx, query)

	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Success insert new customer")
}

func TestQuerySql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	query := "select id, name from customer"
	rows, err := db.QueryContext(ctx, query)
	defer rows.Close()

	if err != nil {
		panic(err.Error())
	}

	//
	for rows.Next() {
		var id, name string
		err := rows.Scan(&id, &name)

		if err != nil {
			panic(err.Error())
		}

		fmt.Println("id:", id)
		fmt.Println("name:", name)
		fmt.Println("----------------------")
	}
}

func TestQuerySqlComplex(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	query := "select id, name, email, balance, rating, birth_date, married, create_at from customer"
	rows, err := db.QueryContext(ctx, query)
	defer rows.Close()

	if err != nil {
		panic(err.Error())
	}

	//
	for rows.Next() {
		var id, name string
		var email sql.NullString // gunakan tipe date jika value null
		var balance int32
		var rating float64
		var createAt time.Time
		var birthdate sql.NullTime // gunakan tipe date jika value null
		var merried bool

		err := rows.Scan(&id, &name, &email, &balance, &rating, &birthdate, &merried, &createAt)

		if err != nil {
			panic(err.Error())
		}

		fmt.Println("id:", id)
		fmt.Println("name:", name)
		if email.Valid {
			fmt.Println("email:", email.String)
		}
		fmt.Println("balnce:", balance)
		fmt.Println("rating:", rating)
		if birthdate.Valid {
			fmt.Println("Birth_date:", birthdate.Time)
		}
		fmt.Println("married:", merried)
		fmt.Println("Create_at:", createAt)
		fmt.Println("----------------------")
	}
}

func TestSqlInjection(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "syarif'; #"
	password := "syarif"

	query := "SELECT username FROM user WHERE username ='" + username + "' AND password='" + password + "' LIMIT 1"
	rows, err := db.QueryContext(ctx, query)

	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	if rows.Next() {
		var username string
		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Println("Sukses Login", username)
	} else {
		fmt.Println("Gagal Login")
	}

}

func TestSqlInjectionSafe(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "syarif"
	password := "syarif"

	query := "SELECT username FROM user WHERE username =? AND password=? LIMIT 1"

	fmt.Println(query)
	rows, err := db.QueryContext(ctx, query, username, password)

	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	if rows.Next() {
		var username string
		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Println("Sukses Login", username)
	} else {
		fmt.Println("Gagal Login")
	}
}

func TestExecSqlParameter(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "syarif'; DROP TABLE user; #"
	password := "syarif"

	script := "INSERT INTO user(username, password) VALUES(?, ?)"
	_, err := db.ExecContext(ctx, script, username, password)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert new user")
}
