package middleware

import (
	"golang-restfull-api/helper"
	"golang-restfull-api/model/web"
	"net/http"
)

type AuthMiddleware struct {
	Handeler http.Handler
}

func NewAuthMiddleware(handeler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{Handeler: handeler}
}

func (middleware *AuthMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if "RAHASIA" == request.Header.Get("x-api-key") {

		middleware.Handeler.ServeHTTP(writer, request)
	} else {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusUnauthorized)

		webResponse := web.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "UNAUTHORIZED",
		}

		helper.WriteToResponseBody(writer, webResponse)
	}
}
