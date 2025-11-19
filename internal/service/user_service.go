package service

import (
	"errors"

	"github.com/NavaneethWKT/CRUD-Go-lang/internal/model"
	"github.com/NavaneethWKT/CRUD-Go-lang/internal/repository"
)

type UserService struct{
	repository *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{
		repository: repo,
	}
}

// create a new user
func (s *UserService) CreateUser(user model.User) (*model.User, error) {
	if user.IsEmpty() {
		return nil, errors.New("user is empty")
	}
	if !user.IsValid() {
		return nil, errors.New("email is invalid")
	}
	if err := s.repository.InsertUser(user); err != nil {
		return nil, err
	}
	return &user, nil
}

// get all users
func (s *UserService) GetAllUsers() ([]model.User, error) {
	users, err := s.repository.GetAllUsers()
	if err != nil {
		return nil, err
	}
	return users, nil
}

// get a specifc user by id
func (s *UserService) GetUserByID(id string) (*model.User, error) {
	user, err := s.repository.GetUserByID(id)
	if err != nil {
		return nil, errors.New("user not found with id: " + id)
	}
	return &user, nil
}

// update a specifc user by id
func (s *UserService) UpdateUser(id string, user model.User) error {
	if user.Name == "" && user.Email == "" && user.Password == "" {
		return errors.New("at least one field must be provided for update")
	}
	success, err := s.repository.UpdateUser(id, user)
	if err != nil {
		return err
	}
	if !success {
		return errors.New("user not found with id: " + id)
	}
	return nil
}