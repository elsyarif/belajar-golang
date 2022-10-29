package golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJSONMapDecode(t *testing.T) {
	jsonString := `{"id": "p123", "name": "msi", "price": 9000000}`
	jsonBytes := []byte(jsonString)

	var result map[string]interface{}
	json.Unmarshal(jsonBytes, &result)

	fmt.Println(result)
	fmt.Println(result["id"])
	fmt.Println(result["name"])
	fmt.Println(result["price"])
}

func TestJSONMapEncode(t *testing.T) {
	priduct := map[string]interface{}{
		"id":    "p123",
		"name":  "msi",
		"price": 9000000,
	}

	bytes, err := json.Marshal(priduct)
	if err != nil {
		return
	}

	fmt.Println(string(bytes))
}
