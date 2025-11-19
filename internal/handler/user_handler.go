package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/NavaneethWKT/CRUD-Go-lang/internal/model"
	"github.com/NavaneethWKT/CRUD-Go-lang/internal/service"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{service: service}
}

// creat a new user
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
		return
	}
	defer r.Body.Close()
	var user model.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		json.NewEncoder(w).Encode("Invalid JSON")
		return
	}
	createdUser, err := h.service.CreateUser(user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	fmt.Println("User created successfully")
	json.NewEncoder(w).Encode(createdUser)
}
