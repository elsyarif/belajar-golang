package golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

// di golnag json array direpresentasikan dengan slice
func TestArrayJSON(t *testing.T) {
	customer := Customer{
		FirstName: "Syarif",
		LastName:  "Hidayatulloh",
		Hobbies:   []string{"Game", "Running"},
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}

func TestArrayJSONDecode(t *testing.T) {
	jsonString := `{"FirstName":"Syarif","LastName":"Hidayatulloh","Age":0,"Merried":false,"Hobbies":["Game","Running"]}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}
	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.Hobbies)
}

func TestArrayJSONComplex(t *testing.T) {
	customer := Customer{
		FirstName: "Syarif",
		Addresses: []Address{
			{
				City:       "Cilegon",
				Country:    "Indonesia",
				PostalCode: "1234",
			},
			{
				City:       "Ciruas",
				Country:    "Indonesia",
				PostalCode: "1234",
			},
		},
	}

	bytes, _ := json.Marshal(customer)

	fmt.Println(string(bytes))
}

func TestArrayJSONComplexDecode(t *testing.T) {
	jsonString := `{"FirstName":"Syarif","LastName":"","Age":0,"Merried":false,"Hobbies":null,"Addresses":[{"City":"Cilegon","Country":"Indonesia","PostalCode":"1234"},{"City":"Ciruas","Country":"Indonesia","PostalCode":"1234"}]}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}
	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.Addresses)
}

func TestOnlyArrayJSONComplexDecode(t *testing.T) {
	jsonString := `[{"City":"Cilegon","Country":"Indonesia","PostalCode":"1234"},{"City":"Ciruas","Country":"Indonesia","PostalCode":"1234"}]`
	jsonBytes := []byte(jsonString)

	address := &[]Address{}
	err := json.Unmarshal(jsonBytes, address)
	if err != nil {
		panic(err)
	}

	fmt.Println(address)
}

func TestOnlyArrayJSONComplex(t *testing.T) {

	Addresses := []Address{
		{
			City:       "Cilegon",
			Country:    "Indonesia",
			PostalCode: "1234",
		},
		{
			City:       "Ciruas",
			Country:    "Indonesia",
			PostalCode: "1234",
		},
	}

	bytes, _ := json.Marshal(Addresses)

	fmt.Println(string(bytes))
}
