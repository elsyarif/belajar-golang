package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestParams(t *testing.T) {
	router := httprouter.New()

	router.GET("/products/:id", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		id := params.ByName("id")
		text := "Product " + id
		fmt.Fprint(writer, text)
	})

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/products/1", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)
	response := recorder.Result()

	bytes, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Product 1", string(bytes))
}
