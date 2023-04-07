package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

// contract of controller

type CategoryController interface {
	Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) // follows handler function contract
	Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
