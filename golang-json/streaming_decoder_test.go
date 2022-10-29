package golang_json

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestStreamingDecoder(t *testing.T) {
	reader, _ := os.Open("customer.json")

	decoder := json.NewDecoder(reader)

	customer := &Customer{}
	decoder.Decode(customer)

	fmt.Println(customer)
}
