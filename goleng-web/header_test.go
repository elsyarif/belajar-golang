package goleng_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func RequestHeader(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("content-type")
	fmt.Fprint(w, contentType)
}

func TestRequestHeader(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello", nil)
	request.Header.Add("Content-Type", "application/json")
	recorder := httptest.NewRecorder()

	// Panggil Handler
	RequestHeader(recorder, request)

	// Cek body response
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	bodyString := string(body)

	fmt.Println(bodyString)
}

func ResponseHeader(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("x-powered-by", "Test Response")
	fmt.Fprint(w, "OK")
}

func TestResponseHeader(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello", nil)
	recorder := httptest.NewRecorder()

	// Panggil Handler
	ResponseHeader(recorder, request)

	// Cek body response
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	bodyString := string(body)
	poweredBy := recorder.Header().Get("x-powered-by")

	fmt.Println(bodyString)
	fmt.Println(poweredBy)
}
