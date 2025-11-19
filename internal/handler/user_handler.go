package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/NavaneethWKT/CRUD-Go-lang/internal/jwt"
	"github.com/NavaneethWKT/CRUD-Go-lang/internal/middleware"
	"github.com/NavaneethWKT/CRUD-Go-lang/internal/model"
	"github.com/NavaneethWKT/CRUD-Go-lang/internal/service"
	"github.com/gorilla/mux"
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

// login user and return JWT token
func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
		return
	}
	defer r.Body.Close()
	var loginRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&loginRequest); err != nil {
		json.NewEncoder(w).Encode("Invalid JSON")
		return
	}
	user, err := h.service.Login(loginRequest.Email, loginRequest.Password)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	token, err := jwt.GenerateToken(user.ID.Hex(), user.Email)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Error generating token: " + err.Error())
		return
	}
	response := map[string]string{
		"token": token,
	}
	fmt.Println("User logged in successfully:", user.Email)
	json.NewEncoder(w).Encode(response)
}

// get all users
func (h *UserHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	users, err := h.service.GetAllUsers()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Error fetching users: " + err.Error())
		return
	}
	fmt.Println("Get all users")
	json.NewEncoder(w).Encode(users)
}

// get a specifc user by id
func (h *UserHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	currentUserID := middleware.GetUserID(r.Context())
	
	params := mux.Vars(r)
	targetUserID := params["id"]
	
	if currentUserID != targetUserID {
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode("You can only view your own profile")
		return
	}
	
	user, err := h.service.GetUserByID(targetUserID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	fmt.Println("Get user by ID:", targetUserID)
	json.NewEncoder(w).Encode(user)
}

// update a specifc user by id
func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	currentUserID := middleware.GetUserID(r.Context())
	
	params := mux.Vars(r)
	targetUserID := params["id"]
	
	if currentUserID != targetUserID {
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode("You can only update your own profile")
		return
	}
	
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
	err := h.service.UpdateUser(targetUserID, user)
	if err != nil {
		if err.Error() == "user not found with id: "+targetUserID {
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	fmt.Println("User updated successfully:", targetUserID)
	json.NewEncoder(w).Encode("User updated successfully")
}

// delete a specifc user by id
func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	currentUserID := middleware.GetUserID(r.Context())
	
	params := mux.Vars(r)
	targetUserID := params["id"]
	
	if currentUserID != targetUserID {
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode("You can only delete your own account")
		return
	}
	
	err := h.service.DeleteUser(targetUserID)
	if err != nil {
		if err.Error() == "user not found with id: "+targetUserID {
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	fmt.Println("User deleted successfully:", targetUserID)
	json.NewEncoder(w).Encode("User deleted successfully")
}

// delete all users
func (h *UserHandler) DeleteAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	count, err := h.service.DeleteAllUsers()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Error deleting users: " + err.Error())
		return
	}
	fmt.Println("All users deleted, count:", count)
	json.NewEncoder(w).Encode("All users deleted, count: " + strconv.FormatInt(count, 10))
}
