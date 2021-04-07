package services

import (
	. "go-service/internal/models"
)

type UserService interface {
	GetAll() (*[]User, error)
	Load(id string) (*User, error)
	Insert(user *User) (int64, error)
	Update(user *User) (int64, error)
	Delete(id string) (int64, error)
}
