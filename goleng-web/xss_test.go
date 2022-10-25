package goleng_web

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TemplateAutoEscape(writer http.ResponseWriter, request *http.Request) {
	myTemplate.ExecuteTemplate(writer, "post.gohtml", map[string]interface{}{
		"Title": "Go-lang Template Auto Escape",
		"Body":  "<p> Selamat Belajar Golang</p>",
	})
}

func TestTemplateAutoEscape(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:9000", nil)
	recorder := httptest.NewRecorder()

	TemplateAutoEscape(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TemplateAutoEscapeDisabled(writer http.ResponseWriter, request *http.Request) {
	myTemplate.ExecuteTemplate(writer, "post.gohtml", map[string]interface{}{
		"Title": "Go-lang Template Auto Escape",
		"Body":  template.HTML("<p> Selamat Belajar Golang</p>"),
	})
}

func TestTemplateAutoEscapeDisabled(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:9000", nil)
	recorder := httptest.NewRecorder()

	TemplateAutoEscapeDisabled(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TemplateXSS(writer http.ResponseWriter, request *http.Request) {
	myTemplate.ExecuteTemplate(writer, "post.gohtml", map[string]interface{}{
		"Title": "Go-lang Template Auto Escape",
		"Body":  template.HTML(request.URL.Query().Get("body")),
	})
}

func TestTemplateXSS(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:9000/?body=<script>alert('Anda di hack')</script>", nil)
	recorder := httptest.NewRecorder()

	TemplateXSS(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TestTemplateXXSServer(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:3000",
		Handler: http.HandlerFunc(TemplateXSS),
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}
