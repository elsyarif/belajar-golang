package main

import (
	"fmt"
	"os"
)

// Package os berisikan fungsionalitas untuk mengakses fitur sistem operasi secara independen (bisa digunakan disemua sistem operasi
func main() {
	args := os.Args

	fmt.Println(args)

	name, _ := os.Hostname()

	fmt.Println(name)

	username := os.Getenv("username")
	password := os.Getenv("password")

	fmt.Println("username:", username)
	fmt.Println(password)
}
