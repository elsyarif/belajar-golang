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

func TestRouterPaternNamedParameter(t *testing.T) {
	router := httprouter.New()

	router.GET("/products/:id/items/:itemId", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		id := params.ByName("id")
		itemId := params.ByName("itemId")

		text := "Product " + id + " Item " + itemId
		fmt.Fprint(writer, text)
	})

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/products/1/items/1", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)
	response := recorder.Result()

	bytes, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Product 1 Item 1", string(bytes))
}

func TestRouterPaternCatchAllParameter(t *testing.T) {
	router := httprouter.New()

	router.GET("/images/*image", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		image := params.ByName("image")

		text := "Images : " + image
		fmt.Fprint(writer, text)
	})

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/images/small/profile.png", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)
	response := recorder.Result()

	bytes, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Images : /small/profile.png", string(bytes))
}
