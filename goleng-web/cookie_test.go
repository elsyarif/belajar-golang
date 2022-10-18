package goleng_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SetCookie(writer http.ResponseWriter, request *http.Request) {
	cookie := new(http.Cookie)
	cookie.Name = "x-cookie-token"
	cookie.Value = request.URL.Query().Get("name")
	cookie.Path = "/"

	http.SetCookie(writer, cookie)
	fmt.Fprint(writer, "Cookie success")
}

func GetCookie(writer http.ResponseWriter, request *http.Request) {
	cookie, err := request.Cookie("x-cookie-token")
	if err != nil {
		fmt.Fprint(writer, "No cookie")
	} else {
		fmt.Fprintf(writer, "Hello %s", cookie)
	}
}

func TestCookie(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/set-cookie", SetCookie)
	mux.HandleFunc("/get-cookie", GetCookie)

	server := http.Server{
		Addr:    "localhost:3003",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func TestSetCookies(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/?name=syarif", nil)
	recorder := httptest.NewRecorder()

	SetCookie(recorder, request)

	cookies := recorder.Result().Cookies()

	for _, cookie := range cookies {
		fmt.Printf("Cookies %s:%s \n", cookie.Name, cookie.Value)
	}
}

func TestGetCookies(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/?name=syarif", nil)
	cookie := new(http.Cookie)
	cookie.Name = "x-cookie-token"
	cookie.Value = "namestest"
	request.AddCookie(cookie)
	recorder := httptest.NewRecorder()

	GetCookie(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}
