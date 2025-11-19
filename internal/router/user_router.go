package router

import (
	"github.com/NavaneethWKT/CRUD-Go-lang/internal/handler"
	"github.com/gorilla/mux"
)


func Router(handler *handler.UserHandler) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/users", handler.CreateUser).Methods("POST")

	return router
}