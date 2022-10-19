package goleng_web

import (
	_ "embed"
	"fmt"
	"net/http"
	"testing"
)

func ServeFile(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Query().Get("name") != "" {
		http.ServeFile(writer, request, "./resources/ok.html")
	} else {
		http.ServeFile(writer, request, "./resources/notfound.html")
	}
}

func TestServeFileServer(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:9000",
		Handler: http.HandlerFunc(ServeFile),
	}
	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}

//go:embed resources/ok.html
var resourceOk string

//go:embed resources/notfound.html
var resourceNotfound string

func ServeFileEmbed(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Query().Get("name") != "" {
		fmt.Fprint(writer, resourceOk)
	} else {
		fmt.Fprint(writer, resourceNotfound)
	}
}

func TestServeFileServerEmbed(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:9000",
		Handler: http.HandlerFunc(ServeFileEmbed),
	}
	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}
