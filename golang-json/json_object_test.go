package golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

// JSON Object di golang dipresentasikan dengan tipe data struct
// Diman setiap attribute di JSON Object merupakan attribut di struct

type Address struct {
	City       string
	Country    string
	PostalCode string
}

type Customer struct {
	FirstName string
	LastName  string
	Age       int
	Merried   bool
	Hobbies   []string
	Addresses []Address
}

func TestJSONObject(t *testing.T) {
	customer := Customer{
		FirstName: "syarif",
		LastName:  "Hidayatulloh",
		Age:       30,
		Merried:   true,
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}
