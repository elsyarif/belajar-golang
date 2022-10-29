package golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Product struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Price    int64  `json:"price"`
	ImageUrl string `json:"image_url"`
}

func TestJSONTag(t *testing.T) {
	product := Product{
		Id:       "001",
		Name:     "MSI",
		Price:    8920128,
		ImageUrl: "http:///example.com/image.ong",
	}

	bytes, _ := json.Marshal(product)
	fmt.Println(string(bytes))
}

func TestJSONTagDecode(t *testing.T) {
	jsonString := `{"id":"001","name":"MSI","price":8920128,"image_url":"http:///example.com/image.ong"}`
	jsonBytes := []byte(jsonString)

	product := &Product{}
	err := json.Unmarshal(jsonBytes, product)
	if err != nil {
		return
	}

	fmt.Println(product)
}
