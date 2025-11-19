package model

import (
	"regexp"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password,omitempty"`
}

func (u *User) IsEmpty() bool {
	return u.Name == "" && u.Email == "" && u.Password == ""
}

func (u *User) IsValid() bool {
	return regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`).MatchString(u.Email)
}