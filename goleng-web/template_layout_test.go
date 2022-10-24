package goleng_web

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TemplateLayout(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles(
		"./template/header.gohtml",
		"./template/footer.gohtml",
		"./template/layout.gohtml",
	))

	t.ExecuteTemplate(writer, "layout", map[string]interface{}{
		"Title": "Tamplate Layout",
		"Name":  "Syarif",
	})
}

func TestTemplateLayout(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:9000", nil)
	recorder := httptest.NewRecorder()

	TemplateLayout(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}
