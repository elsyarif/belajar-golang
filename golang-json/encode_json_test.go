package golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func logJson(data interface{}) {
	bytss, err := json.Marshal(data)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytss))
}

func TestEncode(t *testing.T) {
	logJson("Syarif")
	logJson(1)
	logJson(true)
	logJson([]string{"Syarif", "Hidayatulloh"})
}
