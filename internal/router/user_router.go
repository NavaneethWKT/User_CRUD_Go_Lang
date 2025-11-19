package router

import (
	"github.com/NavaneethWKT/CRUD-Go-lang/internal/handler"
	"github.com/gorilla/mux"
)


func Router(handler *handler.UserHandler) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/users", handler.CreateUser).Methods("POST")
	router.HandleFunc("/api/users/all", handler.GetAllUsers).Methods("GET")
	router.HandleFunc("/api/users/{id}", handler.GetUserByID).Methods("GET")
	router.HandleFunc("/api/users/{id}", handler.UpdateUser).Methods("PUT")

	return router
}