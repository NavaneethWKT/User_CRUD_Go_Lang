package main

import (
	"log"
	"net/http"

	"github.com/NavaneethWKT/CRUD-Go-lang/internal/handler"
	"github.com/NavaneethWKT/CRUD-Go-lang/internal/repository"
	"github.com/NavaneethWKT/CRUD-Go-lang/internal/router"
	"github.com/NavaneethWKT/CRUD-Go-lang/internal/service"
)

func main(){

	repo := repository.NewUserRepository()
	service := service.NewUserService(repo)
	handler := handler.NewUserHandler(service)

	router := router.Router(handler)

	log.Fatal(http.ListenAndServe(":8080", router))
}
