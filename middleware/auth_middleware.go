package middleware

import (
	"go-simple-restful-api/helper"
	"go-simple-restful-api/model/web"
	"net/http"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{Handler: handler}
}

// implementation contract from Handler  in server

func (middleware *AuthMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {

	// check header contains X-API-KEY
	if "AUTH-SECRET" == request.Header.Get("X-API-KEY") {
		// authenticated

		// forward to next Handler
		middleware.Handler.ServeHTTP(writer, request)
	} else {
		// unauthorized

		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusUnauthorized)

		webResponse := web.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "UNAUTHORIZED",
		}

		helper.WriteToResponseBody(writer, webResponse)
	}
}
