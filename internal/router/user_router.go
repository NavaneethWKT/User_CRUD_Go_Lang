package router

import (
	"github.com/NavaneethWKT/CRUD-Go-lang/internal/handler"
	"github.com/NavaneethWKT/CRUD-Go-lang/internal/middleware"
	"github.com/gorilla/mux"
)

func Router(handler *handler.UserHandler) *mux.Router {
	router := mux.NewRouter()

	// Public routes
	router.HandleFunc("/api/users", handler.CreateUser).Methods("POST")
	router.HandleFunc("/api/login", handler.Login).Methods("POST")

	// Protected routes
	router.HandleFunc("/api/users/{id}", middleware.AuthMiddleware(handler.GetUserByID)).Methods("GET")
	router.HandleFunc("/api/users/{id}", middleware.AuthMiddleware(handler.UpdateUser)).Methods("PUT")
	router.HandleFunc("/api/users/{id}", middleware.AuthMiddleware(handler.DeleteUser)).Methods("DELETE")
	router.HandleFunc("/api/users/all", middleware.AuthMiddleware(handler.DeleteAllUsers)).Methods("DELETE")

	return router
}