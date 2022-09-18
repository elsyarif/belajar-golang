package main

import (
	"flag"
	"fmt"
)

// Package flag berisikan fungsionalitas untuk memparsing command line argument
// https://golang.org/pkg/flag/
func main() {
	var host *string = flag.String("host", "localhost", "put yout localhost")
	var username *string = flag.String("username", "root", "put your username")
	var password *string = flag.String("password", "root", "put your password")
	flag.Parse()

	fmt.Println("Host", *host)
	fmt.Println("Username", *username)
	fmt.Println("Password", *password)

	// untuk penggunan-nya pada terminal ketikan
	// go run flag.go -host=localhost -username=root -password=root
}
