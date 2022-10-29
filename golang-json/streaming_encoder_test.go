package golang_json

import (
	"encoding/json"
	"os"
	"testing"
)

// digunakan untuk menulis langsung JSON nya ke IO.Writer
func TestStreamingEncoder(t *testing.T) {
	writer, _ := os.Create("product.json")
	encoder := json.NewEncoder(writer)

	product := Product{
		Id:   "poo12",
		Name: "MSI Modern",
	}

	encoder.Encode(product)
}
